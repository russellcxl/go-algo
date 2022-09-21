package main

import (
	"fmt"
	"time"
)

type nums struct {
	nums []int
}

func main() {

	fmt.Println("TEST")

	arr := []int{1, 2, 3}
	n := nums{nums: arr}
	test3(n)
	fmt.Println(n)
}

func test3(n nums) {
	n.nums = []int{4, 5, 6}
}

func test2() {
	c := make(chan bool, 1)

	rcvr := func() {
		fmt.Println(<-c)
	}

	for i := 0; i < 3; i++ {
		go rcvr()
	}

	for {
		select {
		case c <- true:
			fmt.Println("message sent")
			time.Sleep(time.Second)
		default:
			fmt.Println("message not sent")
			close(c)
			return
		}
	}
}

func test1() {
	rcvr := make(map[int]chan bool)
	rcvr[1] = make(chan bool, 1)
	need := make(chan bool)

	go func() {
		select {
		case <-rcvr[1]:
			fmt.Println("message received")
		case <-time.After(time.Second * 3):
			fmt.Println("timed out")
		}
	}()

	go func() {
		select {
		case <-rcvr[1]:
			fmt.Println("message received")
		case <-time.After(time.Second * 3):
			fmt.Println("timed out")
		}
	}()

	go func() {
		for <-need {
			select {
			case rcvr[1] <- true:
				fmt.Println("message sent")
			default:
				fmt.Println("message not sent")
				return
			}
			time.Sleep(time.Second)
		}
	}()

	need <- true
	need <- true
}
