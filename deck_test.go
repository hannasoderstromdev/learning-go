package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedLength := 52
	length := len(d)

	if length != expectedLength {
		t.Errorf("Expected deck length of " + fmt.Sprintf("%v", expectedLength) + " but got %v", length)
	}

	expectedFirstCard := "Ace of Hearts"
	firstCard := d[0]

	if firstCard != expectedFirstCard {
		t.Errorf("Expected first card to be " + expectedFirstCard + ", but got %s", firstCard)
	}

	expectedLastCard := "King of Clubs"
	lastCard := d[len(d) -1]

	if lastCard != expectedLastCard {
		t.Errorf("Expected last card to be " + expectedLastCard + ", but got %s", lastCard)
	}
}