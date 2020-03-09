package main

import (
	"fmt"
	"time"
)

//var wg sync.WaitGroup

func main() {
	ch := make(chan string, 30)
	//wg.Add(2)
	go timestable(2, ch)
	go timestable(3, ch)
	// Receive values from ch until closed.
	for v := range ch {
		fmt.Println(v)
	}
	//wg.Wait()
}

func timestable(x int, ch chan string) {
	for i := 1; i <= 12; i++ {
		ch <- fmt.Sprintf("%d x %d = "+"%d\n", i, x, x*i)
		time.Sleep(time.Millisecond * 500)
	}
	//wg.Done()
	close(ch)
}
