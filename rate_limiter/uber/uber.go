package uber

import (
	"sync"
	"time"
)

// RateLimiter is the interface implemented by a rate limiter.
type RateLimiter interface {
	// Take should block to make sure that the RPS is met.
	Take()
}

//------------------------------- lock counters --------------------------------

// lockedCounter implements rate limiting using a token bucket.
// The bucket is locked using a mutex, and uses sync.Cond to
// signal waiters.
type lockedCounter struct {
	sync.Mutex
	cond *sync.Cond
	n    int
	rps  int
}

func newLockedCounter(rps int) *lockedCounter {
	c := &lockedCounter{rps: rps}
	c.cond = sync.NewCond(c)
	go c.filler()
	return c
}

func (c *lockedCounter) filler() {
	const numPeriods = 1000

	increment := c.rps / numPeriods
	sleepFor := time.Second / numPeriods

	for i := 0; true; i++ {
		if i == numPeriods {
			i = 0
			c.Lock()
			c.n = increment
			c.Unlock()

		} else {
			c.Lock()
			c.n += increment
			c.Unlock()
		}

		c.cond.Broadcast()
		time.Sleep(sleepFor)
	}
}

func (c *lockedCounter) Take() {
	c.Lock()
	defer c.Unlock()

	if c.n > 0 {
		c.n--
		return
	}

	for c.n <= 0 {
		c.cond.Wait()
	}
	c.n--
}

//------------------------------- time period --------------------------------

type timePeriod struct {
	sync.Mutex
	last       time.Time
	sleepFor   time.Duration
	perRequest time.Duration
}

func newTimePeriod(rps int) *timePeriod {
	return &timePeriod{
		perRequest: time.Second / time.Duration(rps),
	}
}

func (t *timePeriod) Take() {
	t.Lock()
	defer t.Unlock()

	// If this is our first request, then we allow it.
	cur := time.Now()
	if t.last.IsZero() {
		t.last = cur
		return
	}

	// sleepFor calculates how much time we should sleep based on
	// the perRequest budget and how long the last request took.
	// Since the request may take longer than the budget, this number
	// can get negative, and is summed across requests.
	t.sleepFor = t.sleepFor + t.perRequest - cur.Sub(t.last)
	t.last = cur

	// We shouldn't allow sleepFor to get too negative, since it would mean that
	// a service that slowed down a lot for a short period of time would get
	// a much higher RPS following that.
	if t.sleepFor < -time.Second {
		t.sleepFor = time.Second
	}

	// If sleepFor is positive, then we should sleep now.
	if t.sleepFor > 0 {
		time.Sleep(t.sleepFor)
		t.last = cur.Add(t.sleepFor)
		t.sleepFor = 0
	}
}