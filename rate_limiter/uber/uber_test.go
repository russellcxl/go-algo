package uber

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestTimePeriodLimiter(t *testing.T) {
	const numGoRoutines = 5

	l := newTimePeriod(2000)

	var rps int32
	go func() {
		for {
			// print the number of requests per second
			time.Sleep(time.Second)

			// reset rps counter
			v := atomic.SwapInt32(&rps, 0)
			fmt.Println("\nRPS", v)
		}
	}()

	work := func() {
		c := 0
		for {
			l.Take()
			atomic.AddInt32(&rps, 1)
			c++
			if c == 500 {
				fmt.Printf(".")
				c = 0
			}
		}
	}

	for i := 0; i < numGoRoutines; i++ {
		go work()
	}
	select {}
}

func TestLockedCounterLimiter(t *testing.T) {

}
