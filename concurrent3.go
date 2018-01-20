package main

import (
	"fmt"
)

func main() {
	go blah("amit")
}

func blah(msg string) {
	fmt.Println("Blah")
}
