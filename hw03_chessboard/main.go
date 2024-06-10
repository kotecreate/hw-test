package main

import "fmt"

func main() {
	var chessboard, one, two string
	var a uint8
	var err error
	w, b := "  ", "##"
	fmt.Printf("Введите размер доски от 0 до 255: ")
	_, err = fmt.Scanf("%v\n", &a)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	sizeboard := int(a)
	for i := 1; i <= sizeboard; i++ {
		if i%2 != 0 {
			one += w
			two += b
		} else {
			one += b
			two += w
		}
	}
	for i := 1; i <= sizeboard; i++ {
		if i%2 != 0 {
			chessboard += one + "\n"
		} else {
			chessboard += two + "\n"
		}
	}
	fmt.Print(chessboard)
}
