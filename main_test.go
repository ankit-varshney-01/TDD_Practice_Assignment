package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
}

var dir = []string{
	"N",
	"S",
	"E",
	"W",
}

var roverNorth = MarsRover{
	location:  []int{0, 0},
	direction: dir[0],
	gridSize:  []int{50, 50},
}
var roverSouth = MarsRover{
	location:  []int{0, 0},
	direction: dir[1],
	gridSize:  []int{50, 50},
}
var roverEast = MarsRover{
	location:  []int{0, 0},
	direction: dir[2],
	gridSize:  []int{50, 50},
}
var roverWest = MarsRover{
	location:  []int{0, 0},
	direction: dir[3],
	gridSize:  []int{50, 50},
}

func TestInput(t *testing.T) {
	moves := "ffbbff"
	takeInput(&moves)
	if moves == "" {
		t.Errorf("error in taking input of string")
	}
}

func TestInputString(t *testing.T) {
	correctMoves := "ffbbllrrffbb"
	incorrectMoves := "fasdnfsd"

	_, err := validateString(correctMoves)
	require.NoError(t, err)

	_, err = validateString(incorrectMoves)
	require.Error(t, err)
}

func TestForwardMove(t *testing.T) {
	moveForward(&roverNorth)
	if roverNorth.location[0] != 0 && roverNorth.location[1] != 1 {
		t.Errorf("error in moving forward")
	}

	moveForward(&roverSouth)
	if roverSouth.location[0] != 0 && roverSouth.location[1] != -1 {
		t.Errorf("error in moving forward")
	}

	moveForward(&roverEast)
	if roverEast.location[0] != 1 && roverEast.location[1] != 0 {
		t.Errorf("error in moving forward")
	}

	moveForward(&roverWest)
	if roverWest.location[0] != -1 && roverWest.location[1] != 0 {
		t.Errorf("error in moving forward")
	}
}

func TestBackwardMove(t *testing.T) {
	moveBackward(&roverNorth)
	if roverNorth.location[0] != 0 && roverNorth.location[1] != 0 {
		t.Errorf("error in moving forward")
	}

	moveBackward(&roverSouth)
	if roverSouth.location[0] != 0 && roverSouth.location[1] != 0 {
		t.Errorf("error in moving forward")
	}

	moveBackward(&roverEast)
	if roverEast.location[0] != 0 && roverEast.location[1] != 0 {
		t.Errorf("error in moving forward")
	}

	moveBackward(&roverWest)
	if roverWest.location[0] != 0 && roverWest.location[1] != 0 {
		t.Errorf("error in moving forward")
	}
}
