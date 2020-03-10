package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan string)
	for i := 1; i <= 12; i++ {
		wg.Add(1)
		go tables(i, ch)
	}
	// Receive values from ch until closed.
	go display(ch)
	wg.Wait()
	defer close(ch)

}

func tables(x int, ch chan string) {
	for i := 1; i <= 12; i++ {
		ch <- fmt.Sprintf("%d x %d = "+"%d\n", i, x, x*i)
	}
	defer wg.Done()
}

func display(ch chan string){
	for v := range ch {
		fmt.Println(v)
	}
}
