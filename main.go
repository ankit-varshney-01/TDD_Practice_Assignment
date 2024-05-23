package main

import (
	"fmt"
)

func takeInput(moves *string) {
	fmt.Println("Please input order of moves")
	fmt.Scan(moves)
}

func main() {
	var moves string
	takeInput(&moves)
	fmt.Println(moves)
}
