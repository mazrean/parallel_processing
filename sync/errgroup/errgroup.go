package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func despair(ctx context.Context) {
	for {
		fmt.Println("Help me!")

		select {
		case _ = <-ctx.Done():
			fmt.Println("...")
			return
		default:
		}
	}
}

func main() {
	ctx := context.Background()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func(ctx context.Context) func() error {
		return func() error {
			despair(ctx)
			return nil
		}
	}(ctx))

	eg.Go(func() error {
		time.Sleep(1 * time.Second)

		return errors.New("error!")
	})

	err := eg.Wait()
	if err != nil {
		time.Sleep(1 * time.Second)
		panic(err)
	}
}
