package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var dir = []string{
	"N",
	"S",
	"E",
	"W",
}

type MainTestSuite struct {
	suite.Suite

	marsRover MarsRover
}

func (suite *MainTestSuite) SetUpTest() {
	suite.marsRover = MarsRover{
		location:  []int{0, 0},
		direction: dir[1],
		gridSize:  []int{50, 50},
	}
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

func (suite *MainTestSuite) TestInput() {
	t := suite.T()

	t.Run("when input is empty", func(t *testing.T) {
		suite.SetUpTest()

		moves := "ffbbff"
		err := takeInput(&moves)

		require.NoError(t, err)
		// assert.Equal(t, err, inpErr)
	})

	t.Run("when incorrect input string", func(t *testing.T) {
		suite.SetUpTest()

		correctMoves := "ffbbllrrffbb"
		incorrectMoves := "fasdnfsd"

		_, err := validateString(correctMoves)
		require.NoError(t, err)

		_, err = validateString(incorrectMoves)
		require.Error(t, err)
	})
}
func (suite *MainTestSuite) TestForwardMove() {
	t := suite.T()

	t.Run("when moving forward in south dirn", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{1, 0}
		actLoc, err := executeMoves(&suite.marsRover, "f")

		assert.Equal(t, expLoc, actLoc)
		require.NoError(t, err)
	})
	t.Run("when moving forward in south dirn - multiple forward", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{2, 0}
		actLoc, err := executeMoves(&suite.marsRover, "ff")

		assert.Equal(t, expLoc, actLoc)
		require.NoError(t, err)
	})
	t.Run("when moving forward in north dirn", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{49, 0}
		suite.marsRover.direction = dir[0]
		actLoc, err := executeMoves(&suite.marsRover, "f")

		assert.Equal(t, expLoc, actLoc)
		require.NoError(t, err)
	})
	t.Run("when moving forward in east dirn", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{0, 1}
		suite.marsRover.direction = dir[2]
		actLoc, err := executeMoves(&suite.marsRover, "f")

		assert.Equal(t, expLoc, actLoc)
		require.NoError(t, err)
	})
	t.Run("when moving forward in east dirn", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{0, 49}
		suite.marsRover.direction = dir[3]
		actLoc, err := executeMoves(&suite.marsRover, "f")

		assert.Equal(t, expLoc, actLoc)
		require.NoError(t, err)
	})
}
