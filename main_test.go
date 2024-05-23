package main

import "testing"

func TestInput(t *testing.T) {
	moves := "ffbbff"
	takeInput(&moves)
	if moves == "" {
		t.Errorf("error in taking input of string")
	}
}
