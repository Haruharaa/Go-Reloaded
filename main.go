package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide input and output file paths as arguments.")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	err = os.WriteFile(os.Args[2], []byte(Process(string(data))), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
