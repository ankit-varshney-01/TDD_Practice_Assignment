package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInput(t *testing.T) {
	moves := "ffbbff"
	takeInput(&moves)
	if moves == "" {
		t.Errorf("error in taking input of string")
	}
}

func TestInputString(t *testing.T) {
	correctMoves := "ffbbfff"
	incorrectMoves := "fasdnfsd"

	_, err := validateString(correctMoves)
	require.NoError(t, err)

	_, err = validateString(incorrectMoves)
	require.Error(t, err)
}
