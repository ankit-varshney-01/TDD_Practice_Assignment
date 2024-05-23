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

func TestUpMove(t *testing.T) {
	moveUp(roverNorth)
	if roverNorth.location[0] != 0 && roverNorth.location[1] != 1 {
		t.Errorf("error in moving forward")
	}

	moveUp(roverSouth)
	if roverSouth.location[0] != 0 && roverSouth.location[1] != -1 {
		t.Errorf("error in moving forward")
	}

	moveUp(roverEast)
	ifroverEast.location[0] != 1 &&roverEast.location[1] != 0 {
		t.Errorf("error in moving forward")
	}

	moveUp(roverWest)
	if roverWest.location[0] != -1 && roverWest.location[1] != 0 {
		t.Errorf("error in moving forward")
	}
}
