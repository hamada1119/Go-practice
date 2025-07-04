package main

import (
	"fmt"
	"os"
)

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")

}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	TestDefer()

	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	RunDefer()

	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() //ファイルリソースの確実な解放と書き込みの完了を保証

	file.Write([]byte("Hello"))

}
