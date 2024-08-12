package hw03

// func main() {
// 	var a uint8
// 	var err error
// 	fmt.Printf("Введите размер доски от 0 до 255: ")
// 	_, err = fmt.Scanf("%v\n", &a)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 	}
// 	fmt.Print(board(a))
// }

func board(s uint8) string {
	var chessboard, one, two string
	w, b := "  ", "##"
	sizeboard := int(s)
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
	return chessboard
}
