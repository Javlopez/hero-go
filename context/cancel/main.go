package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context err:\t", ctx.Err())
	fmt.Printf("Context type:\t%T\n", ctx)
	fmt.Println("----------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context err:\t", ctx.Err())
	fmt.Printf("Context type:\t%T\n", ctx)
	fmt.Println("Cancel:\t\t", cancel)
	fmt.Printf("Cancel Type:\t%T\n", cancel)
	fmt.Println("----------")

	cancel()

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context err:\t", ctx.Err())
	fmt.Printf("Context type:\t%T\n", ctx)
	fmt.Println("Cancel:\t\t", cancel)
	fmt.Printf("Cancel Type:\t%T\n", cancel)
	fmt.Println("----------")
}
