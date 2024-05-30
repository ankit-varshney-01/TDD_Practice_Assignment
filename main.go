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

func validateString(moves string) (err error) {
	for i := 0; i < len(moves); i = i + 1 {
		if (moves[i] != 'f') && (moves[i] != 'b') && (moves[i] != 'l') && (moves[i] != 'r') {
			return errors.New("invalid character/s: validation failed")
		}
	}

	return nil
}

func takeObstaclesInput(obs *[][]int, noOfObstacles int) (err error) {
	fmt.Println("Please input obstacles locations as space seperated co-ordinates")

	var x, y int
	for i := 0; i < noOfObstacles; i = i + 1 {
		fmt.Scan(&x, &y)
		*obs = append(*obs, []int{x, y})
	}
	fmt.Scan(obs)
	if len(*obs) == 0 {
		return errors.New("obstacles array cannot be empty")
	}

	return nil
}

func validateObstacles(obs [][]int, marsRover *MarsRover) (err error) {
	for i := 0; i < len(obs); i = i + 1 {
		x, y := obs[i][0], obs[i][1]
		fmt.Printf("x: %d, y: %d\n", x, y)
		if x < 0 || x >= marsRover.gridSize[0] || y < 0 || y >= marsRover.gridSize[1] {
			return errors.New("obstacle out of bounds")
		}
	}

	return nil
}

func takeInputObsCnt(n *int) (err error) {
	fmt.Println("Please input no of obstacles")
	fmt.Scan(n)

	if *n < 0 {
		return errors.New("obstacles count cannot be less than 0")
	}

	return nil
}

func executeMoves(rover *MarsRover, moves string, obstacles [][]int) (loc []int, obsLoc []int, err error) {
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

		// checking for obstacles
		for j := 0; j < len(obstacles); j = j + 1 {
			if rover.location[0] == obstacles[j][0] && rover.location[1] == obstacles[j][1] {
				if moves[i] == 'f' {
					moveBackward(rover)
				}
				if moves[i] == 'b' {
					moveForward(rover)
				}
				if moves[i] == 'l' {
					moveRight(&rover.direction)
				}
				if moves[i] == 'r' {
					moveLeft(&rover.direction)
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

				return rover.location, obstacles[j], errors.New("obstacle encountered, returning last position")
			}
		}
	}
	return rover.location, nil, nil
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
	var n int
	var obs [][]int

	var errDetails []error

	err := takeInput(&moves)
	if err != nil {
		errDetails = append(errDetails, err)
	}

	err = takeInputObsCnt(&n)
	if err != nil {
		errDetails = append(errDetails, err)
	}

	err = takeObstaclesInput(&obs, n)
	if err != nil {
		errDetails = append(errDetails, err)
	}

	if len(errDetails) > 0 {
		fmt.Print(errDetails)
	} else {
		fmt.Println("Your input string: ", moves)
		fmt.Println("Your obstacles array: ", obs)

		rover := NewMarsRover([]int{0, 0}, "N", []int{50, 50})

		err = validateString(moves)
		if err != nil {
			errDetails = append(errDetails, err)
		}

		err = validateObstacles(obs, rover)
		if err != nil {
			errDetails = append(errDetails, err)
		}

		if len(errDetails) > 0 {
			fmt.Print(errDetails)
		} else {
			res, obsLoc, err := executeMoves(rover, moves, obs)
			fmt.Println("executed moves")

			if err != nil {
				fmt.Printf("Obstacle encountered at: %v\n", obsLoc)
			}
			fmt.Printf("Final Location is: %v\n", res)
			fmt.Printf("Final Direction is: %s\n", rover.direction)
		}
	}
}
