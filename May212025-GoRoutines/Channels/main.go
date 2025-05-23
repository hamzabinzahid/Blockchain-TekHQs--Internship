package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	myCh := make(chan int, 2)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(ch)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 6
		wg.Done()
	}(myCh, wg)
}
