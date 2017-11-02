package goroutines_demo_test

import (
	"fmt"
	"testing"
	"time"
)

func print100Lines(done chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Print(i)
		time.Sleep(time.Second)
	}

	done <- true
}

func printString(done chan bool) {
	for _, c := range "abcd" {
		fmt.Print(string(c))
		time.Sleep(time.Second * 2)
	}

	done <- true
}

func Test(t *testing.T) {
	doneFirst := make(chan bool)
	doneSecond := make(chan bool)
	go print100Lines(doneFirst)
	go printString(doneSecond)

	<- doneFirst
	<- doneSecond
	
	println("Finish")
}

func TestChannelRange(t *testing.T) {
	channel := make(chan string)
	
	go func(ch chan<- string) {
		for _, c:= range "12345" {
			ch <- string(c)
		}

		close(ch)
	}(channel)

	for v := range channel {
		print(v)
	}
}
