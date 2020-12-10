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
	group, errCtx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		return receiveSignal(errCtx)
	})

	group.Go(func() error {
		return startService(errCtx, &http.Server{Addr: ":9001"})
	})

	group.Go(func() error {
		return startService(errCtx, &http.Server{Addr: ":9002"})
	})

	if err := group.Wait(); err != nil {
		fmt.Printf(err.Error())
	}
}

func receiveSignal(ctx context.Context) error {
	signch := make(chan os.Signal)
	defer close(signch)
	signal.Notify(signch, syscall.SIGINT, syscall.SIGKILL, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTERM)
	select {
	case s := <-signch:
		return fmt.Errorf("Terminated by signal %s", s)
	case <-ctx.Done():
		fmt.Printf("Close signal monitor\n")
		return nil
	}
}

func startService(ctx context.Context, server *http.Server) error {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("Stop service %s\n", server.Addr)
			server.Shutdown(context.TODO())
		}
	}()
	fmt.Printf("Start service %s\n", server.Addr)
	return server.ListenAndServe()
}
