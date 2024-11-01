package main

import (
	"fmt"
	"sync"
)

var count = 0

func r1(m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	count++
	fmt.Printf("Счетчик 1: %d\n", count)
	m.Unlock()
}

func r2(m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	count += 2
	fmt.Printf("Счетчик 2: %d\n", count)
	m.Unlock()
}

func r3(m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	count += 3
	fmt.Printf("Счетчик 3: %d\n", count)
	m.Unlock()
}

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(3)
		go r1(&mutex, &wg)
		go r2(&mutex, &wg)
		go r3(&mutex, &wg)
	}
	wg.Wait()
	fmt.Println("Итог: ", count)
}
