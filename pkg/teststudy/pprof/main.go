package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func logicCode(){
	var c chan int
	for {
		select {
		case v:=<-c:
			fmt.Printf("recv from chan,value:%v\n",v)
		default:
			time.Sleep(500*time.Millisecond)
		}
	}
}

func main(){
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof,"cpu",false,"turn cpu pprof on")
	flag.BoolVar(&isMemPprof,"mem",false,"turn mem pprof on")

	flag.Parse()

	if isCPUPprof{
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed,err:%v\n",err)
			return
		}
		pprof.StartCPUProfile(file)
		defer func() {
			pprof.StopCPUProfile()
			file.Close()
		}()
	}

	for i:=0;i<8;i++{
		go logicCode()
	}
	time.Sleep(20*time.Second)

	if isMemPprof{
		file1,err := os.Create("./mem.pprof")
		if err != nil{
			fmt.Printf("create mem pprof failed,err:%v\n",err)
			return
		}
		pprof.WriteHeapProfile(file1)
		defer file1.Close()
	}
}