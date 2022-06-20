package main

import (
	"fmt"
	"sync"
	"time"
)

func intervalLimiter() {
	in := make(chan bool)

	t := time.NewTicker(time.Second)

	go func() {
		for {
			<-in
			fmt.Printf("task received: %s\n", time.Now().String())

			// blocks until next tick
			<-t.C
		}
	}()

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			in <- true
			defer wg.Done()
		}()
	}

	wg.Wait()
}

func main() {
	intervalLimiter()
}
