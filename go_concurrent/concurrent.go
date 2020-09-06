package go_concurrent

import (
	"fmt"
	"sync"
	"time"
)

var intChannel chan int = make(chan int, 10)
var timeout chan bool = make(chan bool)

var WG sync.WaitGroup

func Loop() {
	for i:=1; i< 11; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 20)
	}

}

func Send() {
	time.Sleep(time.Second * 1)
	intChannel <- 1
	timeout <- false
	time.Sleep(time.Second * 1)
	intChannel <- 2
	timeout <- false
	time.Sleep(time.Second * 1)
	intChannel <- 3
	timeout <- false
	time.Sleep(time.Second * 1)
	timeout <- true

}

func Recv() {
	//num := <- intChannel
	//fmt.Println("Recv: ", num)
	//num = <-intChannel
	//fmt.Println("Recv: ", num)
	//num = <-intChannel
	//fmt.Println("Recv: ", num)

	for {
		select {
			case num := <- intChannel:
				fmt.Println("Recv num:", num)
			case isTimeout:= <- timeout:

				fmt.Println("Recv timeout:", timeout)
				if isTimeout {
					fmt.Println("timeout... ")
				}
		}
	}
}


func Read() {
	for i := 0; i < 5; i++ {
		WG.Add(1)
	}
}

func Write() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("Done ", i)
		WG.Done()
	}
}
