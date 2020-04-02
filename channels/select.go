package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := make(chan int)
	for i := 0; i<5; i++ {
		go sleepyGopher(i, c)
	}

	timeout := time.After(2 * time.Second)
	for i := 0; i<5; i++ {
		select {
			case gopherID := <-c:
				fmt.Println("gopher ", gopherID, " has finished sleeping")
			case <- timeout:
				fmt.Println("too slow")
				return
		}
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000))*time.Millisecond)
	c <- id
}
