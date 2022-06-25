package race_conditions

import (
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

func race() []int {
	var numbers []int
	go appendOne(&numbers)
	go appendTwo(&numbers)
	go appendThree(&numbers)
	time.Sleep(30 * time.Millisecond)
	return numbers
}

func appendThree(numbers *[]int) {
	mu.Lock()
	defer mu.Unlock()
	*numbers = append(*numbers, []int{4, 5, 6}...)
}

func appendTwo(numbers *[]int) {
	mu.Lock()
	defer mu.Unlock()
	*numbers = append(*numbers, []int{2, 3}...)
}

func appendOne(numbers *[]int) {
	mu.Lock()
	defer mu.Unlock()
	*numbers = append(*numbers, 1)
}


