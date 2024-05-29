package main

import (
	"errors"
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

		err := validateString(correctMoves)
		require.NoError(t, err)

		err = validateString(incorrectMoves)
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

func (suite *MainTestSuite) TestBackwardMove() {
	t := suite.T()

	t.Run("when moving backward facing north dirn", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{1, 0}
		suite.marsRover.direction = dir[0]
		actLoc, err := executeMoves(&suite.marsRover, "b")

		assert.Equal(t, expLoc, actLoc)
		require.NoError(t, err)
	})
	t.Run("when moving backward facing north - multiple backward", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{2, 0}
		suite.marsRover.direction = dir[0]
		actLoc, err := executeMoves(&suite.marsRover, "bb")

		assert.Equal(t, expLoc, actLoc)
		require.NoError(t, err)
	})
	t.Run("when moving backward facing south dirn", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{49, 0}
		suite.marsRover.direction = dir[1]
		actLoc, err := executeMoves(&suite.marsRover, "b")

		assert.Equal(t, expLoc, actLoc)
		require.NoError(t, err)
	})
	t.Run("when moving backward facing east dirn", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{0, 49}
		suite.marsRover.direction = dir[2]
		actLoc, err := executeMoves(&suite.marsRover, "b")

		assert.Equal(t, expLoc, actLoc)
		require.NoError(t, err)
	})
	t.Run("when moving backward facing west dirn", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{0, 1}
		suite.marsRover.direction = dir[3]
		actLoc, err := executeMoves(&suite.marsRover, "b")

		assert.Equal(t, expLoc, actLoc)
		require.NoError(t, err)
	})
	t.Run("when moving backward and forward", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{1, 0}
		suite.marsRover.direction = dir[0]
		actLoc, err := executeMoves(&suite.marsRover, "fffbbbb")

		assert.Equal(t, expLoc, actLoc)
		require.NoError(t, err)
	})
}

func (suite *MainTestSuite) TestLeftMove() {
	t := suite.T()

	t.Run("when moving left facing north", func(t *testing.T) {
		suite.SetUpTest()

		expDir := dir[3]
		suite.marsRover.direction = dir[0]
		_, err := executeMoves(&suite.marsRover, "l")
		actDir := suite.marsRover.direction

		assert.Equal(t, expDir, actDir)
		require.NoError(t, err)
	})
	t.Run("when moving left twice facing north", func(t *testing.T) {
		suite.SetUpTest()

		expDir := dir[1]
		suite.marsRover.direction = dir[0]
		_, err := executeMoves(&suite.marsRover, "ll")
		actDir := suite.marsRover.direction

		assert.Equal(t, expDir, actDir)
		require.NoError(t, err)
	})
	t.Run("when moving left facing south", func(t *testing.T) {
		suite.SetUpTest()

		expDir := dir[2]
		_, err := executeMoves(&suite.marsRover, "l")
		actDir := suite.marsRover.direction

		assert.Equal(t, expDir, actDir)
		require.NoError(t, err)
	})
	t.Run("when moving left facing east", func(t *testing.T) {
		suite.SetUpTest()

		expDir := dir[0]
		suite.marsRover.direction = dir[2]
		_, err := executeMoves(&suite.marsRover, "l")
		actDir := suite.marsRover.direction

		assert.Equal(t, expDir, actDir)
		require.NoError(t, err)
	})
	t.Run("when moving left facing west", func(t *testing.T) {
		suite.SetUpTest()

		expDir := dir[1]
		suite.marsRover.direction = dir[3]
		_, err := executeMoves(&suite.marsRover, "l")
		actDir := suite.marsRover.direction

		assert.Equal(t, expDir, actDir)
		require.NoError(t, err)
	})
	t.Run("when moving backward, forward and left", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{48, 1}
		suite.marsRover.direction = dir[0]
		actLoc, err := executeMoves(&suite.marsRover, "fflb")

		assert.Equal(t, expLoc, actLoc)
		require.NoError(t, err)
	})
}

func (suite *MainTestSuite) TestRightMove() {
	t := suite.T()

	t.Run("when moving right facing north", func(t *testing.T) {
		suite.SetUpTest()

		expDir := dir[2]
		suite.marsRover.direction = dir[0]
		_, err := executeMoves(&suite.marsRover, "r")
		actDir := suite.marsRover.direction

		assert.Equal(t, expDir, actDir)
		require.NoError(t, err)
	})
	t.Run("when moving right twice facing north", func(t *testing.T) {
		suite.SetUpTest()

		expDir := dir[1]
		suite.marsRover.direction = dir[0]
		_, err := executeMoves(&suite.marsRover, "rr")
		actDir := suite.marsRover.direction

		assert.Equal(t, expDir, actDir)
		require.NoError(t, err)
	})
	t.Run("when moving right facing south", func(t *testing.T) {
		suite.SetUpTest()

		expDir := dir[3]
		_, err := executeMoves(&suite.marsRover, "r")
		actDir := suite.marsRover.direction

		assert.Equal(t, expDir, actDir)
		require.NoError(t, err)
	})
	t.Run("when moving right facing east", func(t *testing.T) {
		suite.SetUpTest()

		expDir := dir[1]
		suite.marsRover.direction = dir[2]
		_, err := executeMoves(&suite.marsRover, "r")
		actDir := suite.marsRover.direction

		assert.Equal(t, expDir, actDir)
		require.NoError(t, err)
	})
	t.Run("when moving right facing west", func(t *testing.T) {
		suite.SetUpTest()

		expDir := dir[0]
		suite.marsRover.direction = dir[3]
		_, err := executeMoves(&suite.marsRover, "r")
		actDir := suite.marsRover.direction

		assert.Equal(t, expDir, actDir)
		require.NoError(t, err)
	})
	t.Run("when moving backward, forward, left and right", func(t *testing.T) {
		suite.SetUpTest()

		expLoc := []int{49, 1}
		expdir := dir[3]
		suite.marsRover.direction = dir[0]
		actLoc, err := executeMoves(&suite.marsRover, "fflblfr")
		actDir := suite.marsRover.direction

		assert.Equal(t, expLoc, actLoc)
		assert.Equal(t, expdir, actDir)
		require.NoError(t, err)
	})
}

func (suite *MainTestSuite) TestObstaclesInput() {
	t := suite.T()

	t.Run("when incorrect obstacles array", func(t *testing.T) {
		suite.SetUpTest()

		// current marsRover grid size = [50, 50]
		obs := [][]int{
			{1, 0},
			{4, 5},
			{50, 2},
		}

		obsErr := errors.New("obstacle out of bounds")
		err := validateObstacles(obs, &suite.marsRover)

		assert.Equal(t, obsErr, err)
		require.Error(t, err)
	})
}
