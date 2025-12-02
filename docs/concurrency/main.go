package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Goroutine")
	// simpleGoroutine()
	decoupleRoutine()
	handleError()
	waitGroup()
}

func simpleGoroutine() {
	// make channel to send data to other goroutine
	data := make(chan string)
	go func() {
		fmt.Println("goroutine 1")
		res := <-data
		fmt.Println(res)
	}()

	go func() {
		fmt.Println("goroutine 2")
		data <- "from other routines"
	}()

	time.Sleep(time.Second)
}

func decoupleRoutine() {
	// make channel to send data to other goroutine
	data := make(chan string)
	go firstRoutine(data)
	go secondRoutine(data)
	time.Sleep(time.Second)
}

func firstRoutine(data chan string) {
	fmt.Println("goroutine 1")
	res := <-data
	fmt.Println(res)
}

func secondRoutine(data chan string) {
	fmt.Println("goroutine 2")
	data <- "from other routines"
}

// handle error with goroutine

func handleError() {
	fmt.Println("---Handle Error---")
	data := make(chan int)
	err := make(chan string)

	go func() {
		for i := range 10 {
			data <- i
		}
		err <- "error"
	}()

	go func() {
		for {
			select {
			case res := <-err:
				fmt.Println(res)
				return
			case datas := <-data:
				fmt.Println(datas)
			}
		}
	}()
	time.Sleep(time.Second)
}

// Wait group, wait other goroutine done

func waitGroup() {
	fmt.Println("---Wait Group---")
	wg := new(sync.WaitGroup)
	data := make(chan int)
	err := make(chan string)
	wg.Add(2) // total goroutine

	go func(w *sync.WaitGroup) {
		defer wg.Done()
		for i := range 10 {
			data <- i
		}
		err <- "error"
	}(wg)

	go func(w *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case res := <-err:
				fmt.Println(res)
				return
			case datas := <-data:
				fmt.Println(datas)
			}
		}
	}(wg)
	wg.Wait()
}
