package main

import (
	"fmt"
	"time"
)

const (
	MAX_LIMIT = 3
	REQUESTS  = 10
)

// handles 3 reqs at a time
func burstLimiter() {
	buf := make(chan bool, MAX_LIMIT)

	for i := 0; i < 3; i++ {
		buf <- true
	}

	t := time.NewTicker(time.Second)
	go func() {
		for {
			<-t.C

			// I really don't understand why this has to be MAX_LIMIT - 1
			for len(buf) < MAX_LIMIT - 1 {
				buf <- true
			}
		}
	}()

	// create reqs
	reqs := make(chan bool, REQUESTS)
	for i := 0; i < REQUESTS; i++ {
		reqs <- true
	}
	// if you don't close, program will not stop running
	// in most cases, you don't have to close
	// but since there is only 1 sender, and you know you're not sending any more values
	close(reqs)

	for _ = range reqs {
		<-buf
		fmt.Printf("task received: %s\n", time.Now().String())
	}

}

func main() {
	burstLimiter()
}
