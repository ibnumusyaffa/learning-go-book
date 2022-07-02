package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	//create context
	// ctx := context.TODO()
	ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)

	// go func() {
	// 	s := bufio.NewScanner(os.Stdin)
	// 	s.Scan()
	// 	cancel()
	// }()
	// time.AfterFunc(time.Second, cancel)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

}

func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
