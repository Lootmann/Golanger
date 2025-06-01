package main

import (
	"fmt"
	"os"
)

func readfile_test() {
	fmt.Println("* File Test")
	read, _ := os.ReadFile("sample.txt")
	fmt.Print(string(read))
}

func directory_test() {
	fmt.Println("* Directory Test")
	path, err := os.Getwd()
	fmt.Println(path, err)
}

func main() {
	// readfile_test()
	directory_test()
}
