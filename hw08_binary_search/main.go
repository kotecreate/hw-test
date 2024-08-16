package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	slice := randSli()
	indVal := randVal(slice)
	value := slice[indVal]
	fmt.Printf("Создан массив: %v\n", slice)
	fmt.Printf("Найти: %v\n", value)
	fmt.Printf("Искомое: %v\n", binaryS(slice, value))
}

func randSli() (res []int) { //создание рандомного слайса
	randLen := rand.Intn(100)
	for i := 0; i < randLen; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		res = append(res, r.Int())
	}
	sort.Ints(res)
	return res
}

func binaryS(sl []int, val int) int {
	var valMid int
	var mid int
	left := 0
	right := (len(sl) - 1)
	for left <= right {
		mid = (left + right) / 2
		valMid = sl[mid]
		if valMid == val {
			return sl[mid]
		} else if valMid < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return valMid
}

func randVal(sl []int) int { //создание рандомного индекса в слайсе
	v := rand.Intn(len(sl))
	return v
}
