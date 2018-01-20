package main

import (
	"fmt" 
	"time"
	"math/rand"
)

func main() {
	c := service("service 1")
	c1:= service("service 2")
	for i := 1; i<5; i++ {
		fmt.Printf(<-c)
		fmt.Printf(<-c1)
	}
	fmt.Println("I am going\n");
}

func service(msg string) <-chan string {
	c:=make(chan string)
	go func() {
		for i:=0;;i++ {
			c <- fmt.Sprintf("%s says %d\n", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
