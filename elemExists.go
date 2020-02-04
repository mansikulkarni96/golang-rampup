package main

import "fmt"

func main() {

	slice := []string{"abc", "bhi", "ghi", "jkl"}

	if elemExists("def", slice) {
		fmt.Println("Exists!")
	}else{
		fmt.Println("Not found")
	}
}

func elemExists(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
