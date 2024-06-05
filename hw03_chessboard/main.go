package main

import "fmt"

func main() {
	var chessboard, one, two string
	var a uint8
	w, b := "  ##", "##  "
	fmt.Printf("Введите размер доски: ")
	fmt.Scanf("%v\n", &a)
	sizeboard := int(a)
	for i := 1; i <= sizeboard/2; i++ {
		one += w
		two += b
	}
	for i := 1; i <= sizeboard/2; i++ {
		chessboard = chessboard + one + "\n" + two + "\n"
	}
	fmt.Print(chessboard)
}
