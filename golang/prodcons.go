package main

import (
	"fmt"
	"sync"
)

func producer(c chan int, wg *sync.WaitGroup) {

	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	wg.Done()

}

func consumer(c chan int, s string, wg *sync.WaitGroup) {
	for v := range c {
		fmt.Println(s, v)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	//exit := make(chan bool)
	//sync := make(chan bool)

	wg.Add(3)
	go consumer(ch, "consumer 1", &wg)
	go consumer(ch, "consumer 2", &wg)
	go producer(ch, &wg)

	wg.Wait()

}
