package main

import "fmt"

func main() {
	var chessboard, one, two string
	var a uint8
	w, b := "  ", "##"
	for {
		fmt.Printf("Введите размер доски от 1 до 255: ")
		fmt.Scanf("%v\n", &a)
		if a > 0 && a <= 255 {
			break
		}
		fmt.Println("Неверное значение")
	}
	sizeboard := int(a)
	for i := 1; i <= sizeboard; i++ {
		one += w + b
		two += b + w
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
