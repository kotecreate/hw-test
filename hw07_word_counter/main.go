package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	text := "Привет мир! привет привет!!!"
	fmt.Printf("%v\n", countWords(text))
}

func countWords(str string) map[string]int {
	var err error
	sl := cleanWords(str)
	if err != nil {
		fmt.Println(err)
	}
	cache := map[string]int{}
	for i := 0; i < len(sl); i++ {
		if _, ok := cache[sl[i]]; !ok {
			cache[sl[i]] = 1
		} else {
			cache[sl[i]]++
		}
	}
	return cache
}

func cleanWords(st string) []string {
	var slice []string
	if utf8.ValidString(st); false {
		fmt.Println("Cтрока не валидна!")
	}
	reg := regexp.MustCompile(`[[:punct:]]`)
	st = reg.ReplaceAllString(st, "")
	st = strings.ToLower(st)
	slice = append(slice, strings.Fields(st)...)
	return slice
}
