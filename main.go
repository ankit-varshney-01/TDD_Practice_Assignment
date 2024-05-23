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

func takeInput(moves *string) (err error) {
	fmt.Println("Please input order of moves")
	fmt.Scan(moves)
	if *moves == "" {
		return errors.New("string cannot be empty")
	}

	return nil
}

func validateString(moves string) (validate bool, err error) {
	for i := 0; i < len(moves); i = i + 1 {
		if (moves[i] != 'f') && (moves[i] != 'b') && (moves[i] != 'l') && (moves[i] != 'r') {
			return false, errors.New("invalid character: validation failed")
		}
	}

	return true, nil
}

func executeMoves(rover *MarsRover, moves string) (loc []int, err error) {
	for i := 0; i < len(moves); i = i + 1 {
		if moves[i] == 'f' {
			moveForward(rover)
		}
		if moves[i] == 'b' {
			moveBackward(rover)
		}
		if moves[i] == 'l' {
			moveLeft(&rover.direction)
		}
		if moves[i] == 'r' {
			moveRight(&rover.direction)
		}

		// To handle edge cases for out of bounds move
		if rover.location[0] < 0 {
			rover.location[0] = rover.gridSize[0] - 1
		}

		if rover.location[1] < 0 {
			rover.location[1] = rover.gridSize[1] - 1
		}

		if rover.location[0] >= rover.gridSize[0] {
			rover.location[0] = 0
		}

		if rover.location[1] >= rover.gridSize[1] {
			rover.location[1] = 0
		}
	}
	return rover.location, nil
}

func moveForward(rover *MarsRover) {
	if rover.direction == "N" {
		rover.location[0] = rover.location[0] - 1
	}
	if rover.direction == "S" {
		rover.location[0] = rover.location[0] + 1
	}
	if rover.direction == "E" {
		rover.location[1] = rover.location[1] + 1
	}
	if rover.direction == "W" {
		rover.location[1] = rover.location[1] - 1
	}
}

func moveBackward(rover *MarsRover) {
	if rover.direction == "N" {
		rover.location[0] = rover.location[0] + 1
	}
	if rover.direction == "S" {
		rover.location[0] = rover.location[0] - 1
	}
	if rover.direction == "E" {
		rover.location[1] = rover.location[1] - 1
	}
	if rover.direction == "W" {
		rover.location[1] = rover.location[1] + 1
	}
}

func moveLeft(curDir *string) {
	if *curDir == "N" {
		*curDir = "W"
	} else if *curDir == "S" {
		*curDir = "E"
	} else if *curDir == "E" {
		*curDir = "N"
	} else if *curDir == "W" {
		*curDir = "S"
	}
}

func moveRight(curDir *string) {
	if *curDir == "N" {
		*curDir = "E"
	} else if *curDir == "S" {
		*curDir = "W"
	} else if *curDir == "E" {
		*curDir = "S"
	} else if *curDir == "W" {
		*curDir = "N"
	}
}

func main() {
	var moves string
	err := takeInput(&moves)

	fmt.Println("Your input string: ", moves)

	_, err = validateString(moves)
	fmt.Print(err)
}
