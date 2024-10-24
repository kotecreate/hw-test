package main

import (
	"fmt"
	"testing"
)

func TestSensor(t *testing.T) {
	testSens := make(chan int)
	go sensor(testSens)
	for i := 0; i < 10; i++ {
		val, ok := <-testSens
		if !ok {
			t.Errorf("Channel is closed before expected")
			return
		}
		fmt.Println("Received value: ", val)
	}
}

func TestProcessed(t *testing.T) {
	testSens2 := make(chan int)
	testResult := make(chan int)
	go sensor(testSens2)
	go processed(testSens2, testResult)
	val, ok := <-testResult
	if !ok {
		t.Errorf("Канал передачи данных закрыт")
		return
	}
	fmt.Println("Среднее арифмитическое: ", val)
}
