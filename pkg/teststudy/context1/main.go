package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exitChan = make(chan bool,1)

func f1(){
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("haha")
		time.Sleep(time.Millisecond*500)
		select {
		case <- exitChan:
			break FORLOOP
		default:
		}
	}
}

func main(){
	wg.Add(1)
	go f1()
	time.Sleep(5*time.Second)
	exitChan <- true
	wg.Wait()

}
