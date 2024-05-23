package main

import "testing"

func TestInput(t *testing.T) {
	inputString := takeInput()
	if inputString == nil {
		t.Errorf("error in taking input of string")
	}
}
