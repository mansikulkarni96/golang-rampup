package main

import (
	"fmt"
	"sync"
)
var sg sync.WaitGroup
func main() {
	for n := 2; n <= 12; n++ {
		sg.Add(1)
		go timetable(n)
	}
	sg.Wait()
}

func timetable(x int) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, x, x*i)

	}
	sg.Done()
}
