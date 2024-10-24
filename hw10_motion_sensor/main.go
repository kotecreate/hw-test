package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sensor(sens chan int) {
	timeout := time.After(1 * time.Minute)
	var data int
	for {
		select {
		case <-timeout:
			time.Sleep(1 * time.Second)
			fmt.Println("Timeout")
			close(sens)
			return
		case sens <- data:
			data = rand.Intn(100)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func processed(sens, res chan int) {
	for {
		math := 0
		for i := 0; i < 10; i++ {
			val, ok := <-sens
			if !ok {
				fmt.Println("Channel is closed")
				close(res)
				return
			}
			math += val
		}
		res <- math / 10
	}
}

func main() {
	sens := make(chan int)
	res := make(chan int)
	go sensor(sens)
	go processed(sens, res)
	for {
		val, ok := <-res
		if !ok {
			fmt.Println("Канал передачи данных закрыт")
			return
		}
		fmt.Println("Среднее арифмитическое: ", val)
	}
}
