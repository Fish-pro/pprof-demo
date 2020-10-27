package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func worker(ctx context.Context){
Loop:
	for{
		fmt.Println("db connecting...")
		time.Sleep(time.Millisecond*50)
		select {
		case <- ctx.Done():
			fmt.Println("timeout")
			break Loop
		default:
		}
	}
	fmt.Println("worker done")
	wg.Done()
}

func main(){
	ctx,cancel := context.WithTimeout(context.Background(),time.Microsecond*10)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second*5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
