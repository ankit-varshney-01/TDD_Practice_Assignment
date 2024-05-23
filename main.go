package main

import (
	"errors"
	"fmt"
)

type MarsRover struct {
	location  []int
	direction string
	gridSize  []int
}

func NewMarsRover(loc []int, dir string, gs []int) *MarsRover {
	rover := new(MarsRover)
	rover.location = loc
	rover.direction = dir
	rover.gridSize = gs

	return rover
}

func takeInput(moves *string) {
	fmt.Println("Please input order of moves")
	fmt.Scan(moves)
}

func validateString(moves string) (validate bool, err error) {
	for i := 0; i < len(moves); i = i + 1 {
		if (moves[i] != 'f') && (moves[i] != 'b') && (moves[i] != 'l') && (moves[i] != 'r') {
			return false, errors.New("invalid character: validation failed")
		}
	}

	return true, nil
}

func moveForward(rover *MarsRover) {
	if rover.direction == "N" {
		rover.location[1] = rover.location[1] + 1
	}
	if rover.direction == "S" {
		rover.location[1] = rover.location[1] - 1
	}
	if rover.direction == "E" {
		rover.location[0] = rover.location[0] + 1
	}
	if rover.direction == "W" {
		rover.location[0] = rover.location[0] - 1
	}
}

func main() {
	var moves string
	takeInput(&moves)

	fmt.Println("Your input string: ", moves)

	_, err := validateString(moves)
	fmt.Print(err)
}
