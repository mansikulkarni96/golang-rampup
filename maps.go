package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["a"] = "Learn"
	m["b"] = "go"
	m["c"] = "using"
	m["d"] = "maps"

	for k, v := range m {
		fmt.Printf("The key is %s and the value is %s", k, v)
		fmt.Print("\n")
	}
	_,ok := m["c"]
	fmt.Print(ok)

}