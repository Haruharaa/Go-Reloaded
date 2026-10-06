package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}
	fmt.Println("Entrée :", os.Args[1], "| Sortie :", os.Args[2])
}
