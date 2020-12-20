package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"week04/api/hello/v1"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return receiveSignal(ctx)
	})

	g.Go(func() error {
		return startServer(ctx, ":9000")
	})

	fmt.Println(g.Wait())
}

func receiveSignal(ctx context.Context) error {
	signch := make(chan os.Signal)
	defer close(signch)
	signal.Notify(signch, syscall.SIGINT, syscall.SIGKILL, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTERM)
	select {
	case s := <-signch:
		return fmt.Errorf("Terminate via signal %s", s)
	case <-ctx.Done():
		fmt.Printf("Close signal listener\n")
		return nil
	}
}

func startServer(ctx context.Context, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("Stop server %s\n", address)
			s.GracefulStop()
		}
	}()
	fmt.Printf("Start server %s\n", address)
	hello.RegisterGreeterServer(s, InitializeHelloServer("db address"))
	return s.Serve(lis)
}
