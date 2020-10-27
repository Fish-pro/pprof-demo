package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notify bool

func f1(){
	defer wg.Done()
	for {
		fmt.Println("haha")
		time.Sleep(time.Millisecond*500)
		if notify{
			break
		}
	}
}

func main(){
	wg.Add(1)
	go f1()
	time.Sleep(5*time.Second)
	notify = true
	wg.Wait()

}
