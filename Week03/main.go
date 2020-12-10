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
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return receiveSignal(ctx)
	})

	g.Go(func() error {
		return startServer(ctx, &http.Server{Addr: ":9001"})
	})

	g.Go(func() error {
		return startServer(ctx, &http.Server{Addr: ":9002"})
	})

	fmt.Println(g.Wait())
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

func startServer(ctx context.Context, server *http.Server) error {
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
