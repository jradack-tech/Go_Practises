package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		counter += 1
		go func() {
			defer func() {
				fmt.Println(counter, "goroutines remaining")
				counter -= 1
				wg.Done()
			}()
			duration := time.Duration(rand.Intn(500) * int(time.Millisecond))
			fmt.Println("Waiting for", duration, "Ref:", time.Millisecond)
			time.Sleep(duration)
		}()
	}

	fmt.Println("Waiting for goroutines to finish")
	wg.Wait()
	fmt.Println("Done!")
}
