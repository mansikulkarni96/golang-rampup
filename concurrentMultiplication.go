package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan string, 30)
	wg.Add(12)
	for i := 1; i <= 12; i++ {
		go timestable(i, ch)
	}
	// Receive values from ch until closed.
	go display(ch)
	wg.Wait()
	defer close(ch)

}

func timestable(x int, ch chan string) {
	for i := 1; i <= 12; i++ {
		ch <- fmt.Sprintf("%d x %d = "+"%d\n", i, x, x*i)
		//time.Sleep(time.Millisecond * 500)
	}
	defer wg.Done()
}

func display(ch chan string){
	for v := range ch {
		fmt.Println(v)
	}
}
