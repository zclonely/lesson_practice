package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	for i := 0; i < 10; i++ {
// 		go fmt.Println(i)
// 	}
// 	time.Sleep(8 * time.Second)
// }

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		fmt.Println("input int to ch")
// 		ch <- 0
// 	}()
// 	time.Sleep(2 * time.Second)
// }

func main() {
	ch := make(chan int, 10)
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10)
			fmt.Println("putting number", n)
			ch <- n
		}
	}()
	go func() {
		for {
			n := <-ch

			time.Sleep(time.Second)
			fmt.Println("receiving ", n)
		}
	}()
	for {
		time.Sleep(time.Second)
	}
}
