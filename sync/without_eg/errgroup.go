package main

import (
	"context"
	"errors"
	"fmt"
	"time"
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

	errChan := make(chan error)

	go despair(ctx)

	go func() {
		time.Sleep(1 * time.Second)

		errChan <- errors.New("error!")
		errChan <- nil
	}()

	err := <-errChan
	if err != nil {
		time.Sleep(1 * time.Second)
		panic(err)
	}
}
