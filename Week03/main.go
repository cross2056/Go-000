package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return receiveSignal(errCtx, cancel)
	})

	group.Go(func() error {
		return startService(errCtx, group, cancel, &http.Server{Addr: ":9001"})
	})

	group.Go(func() error {
		return startService(errCtx, group, cancel, &http.Server{Addr: ":9002"})
	})

	if err := group.Wait(); err != nil {
		fmt.Printf(err.Error())
	}
}

func receiveSignal(ctx context.Context, cancel context.CancelFunc) error {
	signch := make(chan os.Signal)
	defer close(signch)
	signal.Notify(signch, syscall.SIGINT, syscall.SIGKILL, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTERM)
	select {
	case s := <-signch:
		fmt.Printf("Got an exit signal %s\n", s)
		cancel()
		return nil
	case <-ctx.Done():
		fmt.Printf("Close signal monitor\n")
		return nil
	}
}

func startService(ctx context.Context, group *errgroup.Group, cancel context.CancelFunc, server *http.Server) error {
	defer cancel()
	group.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Printf("stop service %s\n", server.Addr)
			return server.Shutdown(context.TODO())
		}
	})
	fmt.Printf("start service %s\n", server.Addr)
	return server.ListenAndServe()
}
