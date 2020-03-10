package main

import (
	"fmt"
)

func main() {
	for n := 2; n <= 12; n++ {
		wg.Add(1)
		go timetable(n)
	}
	wg.Wait()
}

func timetable(x int) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, x, x*i)

	}
	wg.Done()
}
