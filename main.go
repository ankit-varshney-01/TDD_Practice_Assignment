package main

import (
	"errors"
	"fmt"
)

func takeInput(moves *string) {
	fmt.Println("Please input order of moves")
	fmt.Scan(moves)
}

func validateString(moves string) (validate bool, err error) {
	for i := 0; i < len(moves); i = i + 1 {
		if moves[i] != 'f' && (moves[i] != 'b') && (moves[i] != 'l') && (moves[i] != 'r') {
			return false, errors.New("invalid character: validation failed")
		}
	}

	return true, nil
}

func main() {
	var moves string
	takeInput(&moves)

	fmt.Println("Your input string: ", moves)

	_, err := validateString(moves)
	fmt.Print(err)
}
