package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	log.Print(ctx.Value("mykey"))

	ctx2 := context.WithValue(ctx, "mykey", 123)
	log.Print(ctx2.Value("mykey"))

	// context cancelled
	ctx3, ctxCancel := context.WithCancel(ctx)
	log.Print(ctx3.Err())
	ctxCancel()
	log.Print(ctx3.Err())

	// Context deadline
	t := time.Now().Add(1 * time.Second)
	ctx4, _ := context.WithDeadline(ctx, t)
	log.Print(ctx4.Err())
	time.Sleep(2 * time.Second)
	log.Print(ctx4.Err())

}
