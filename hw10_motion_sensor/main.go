package main

import (
	"fmt"
	"time"
)

func sensor(sens chan float64) {
	data := 1.0
	for {
		select {
		case sens <- data:
			data += 1.0
			time.Sleep(100 * time.Millisecond)
		default:
			fmt.Println("Channel is full, skipping data")
			close(sens)
			return
		}
	}
}

func processed(sens, res chan float64) {
	timeout := time.After(1 * time.Minute)
	for {
		math := 0.0
		select {
		case <-timeout:
			time.Sleep(1 * time.Second)
			fmt.Println("Timeout")
			close(res)
			return
		default:
			for i := 0; i < 10; i++ {
				val, ok := <-sens
				if !ok {
					fmt.Println("Channel is closed")
					close(res)
					return
				}
				math += val
			}
		}
		res <- math / 10
	}
}

func main() {
	sens := make(chan float64)
	res := make(chan float64)

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
