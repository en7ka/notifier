package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/en7ka/notifier/internal/closer"
	"github.com/en7ka/notifier/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initdeps(ctx); err != nil {
		return nil, err
	}

	if err := a.initServiceProvider(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initdeps(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func gracefulShutdown(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-waitSignal():
		log.Println("terminating: via signal")
	}

	cancel()
	if wg != nil {
		wg.Wait()
	}
}

func waitSignal() chan os.Signal {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	return sigterm
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
	}()

	if a.serviceProvider == nil {
		return fmt.Errorf("service provider is nil")
	}

	ctx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := a.serviceProvider.SenderConsumer().RunConsumer(ctx)
		if err != nil {
			log.Printf("failed to run consumer: %s", err.Error())
		}
	}()

	gracefulShutdown(ctx, cancel, wg)
	return nil
}
