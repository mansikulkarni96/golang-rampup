
package main

import (
	"fmt"
)

func main() {
	type Dog struct {
		Name  string
		Color string
	}
	var King Dog
	King.Name = "King"
	King.Color = "Black"
	fmt.Println(King.Name, King.Color)
}
