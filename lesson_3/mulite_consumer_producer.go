package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(pname string, ch chan<- int, num_range int) {
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(num_range)
		fmt.Println("pname:", pname, "is putting number ", n)
		ch <- n
	}
}

func consumer(cname string, ch <-chan int) {
	for {
		n := <-ch
		time.Sleep(time.Second)
		fmt.Println("cname: ", cname, "is receiving ", n)
	}
}

func main() {
	ch := make(chan int)
	go producer("p1", ch, 100)
	go producer("p2", ch, 1000)
	go consumer("c1", ch)
	go consumer("c2", ch)
	time.Sleep(10 * time.Second)
	// close(ch)
	fmt.Println("closing channel")
}
