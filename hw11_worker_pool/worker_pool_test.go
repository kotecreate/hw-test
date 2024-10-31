package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 200; i++ {
		wg.Add(3)
		go r1(&mutex, &wg)
		go r2(&mutex, &wg)
		go r3(&mutex, &wg)
	}
	wg.Wait()
	if count != 600 {
		t.Error("Неверное значение счетчика")
	} else {
		fmt.Println("Итог: ", count)
	}
}
