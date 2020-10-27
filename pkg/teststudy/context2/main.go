package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exitChan = make(chan bool,1)

func f1(ctx context.Context){
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("haha")
		time.Sleep(time.Millisecond*500)
		select {
		case <- ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func main(){
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f1(ctx)
	time.Sleep(5*time.Second)
	cancel()
	wg.Wait()

}
