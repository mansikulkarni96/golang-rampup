package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan string)
	wg.Add(1)
	go tables(2, ch)
	// Receive values from ch until closed.
	go display(ch)
	wg.Wait()
	defer close(ch)

}

func tables(x int, ch chan string) {
	for i := 1; i <= 10; i++ {
		ch <- fmt.Sprintf("%d x %d = "+"%d\n", x, i, x*i)
	}
	defer wg.Done()
}

func display(ch chan string){
	for v := range ch {
		fmt.Println(v)
	}
}
