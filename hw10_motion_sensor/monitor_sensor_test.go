package main

import (
	"testing"
)

func TestSensor(t *testing.T) {
	testSens := make(chan float64)
	go sensor(testSens)
}

func TestProcessed(t *testing.T) {
	testSens := make(chan float64)
	testResult := make(chan float64)
	go sensor(testSens)
	go processed(testSens, testResult)
}
