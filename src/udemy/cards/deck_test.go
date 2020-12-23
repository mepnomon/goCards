package main

import (
	"errors"
	"os"
	"testing"
)

func Test_createNewDeck_success(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected lengh of 16, but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" && d[len(d)-1] != "Ace of Diamonds" {
		t.Errorf("Test Failed.")
	}
}

func Test_saveToDeck_success(t *testing.T) {
	os.Remove("test_deck")
	d := newDeck()
	err := d.saveToFile("test_deck")
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}
}

func Test_newDeckFromFile_success(t *testing.T) {
	_, err := loadDeckFromFile("test_deck")
	if err != nil {
		t.Errorf("Test failed, %v", err)
	}
	os.Remove("test_deck")
}

func Test_saveToDeck_and_newDeckFromFile(t *testing.T) {
	os.Remove("_deck_testing")
	d := newDeck()
	d.saveToFile("_deck_testing")

	loadedDeck, err := loadDeckFromFile("_deck_testing")
	if len(loadedDeck) != 16 {
		t.Errorf("Test failed: Expected 16, got  %v . Err: %v", len(loadedDeck), err)
	}
	err = os.Remove("_deck_testing")
	if err != nil {
		myError := errors.New("\"Overly Pedantic Stack Overflow Exception\"")
		t.Errorf("Test failed: %v caused by %v", myError, err)
	}
}

func Test_testShuffle_success(t *testing.T) {
	d := newDeck()
	dCopy := newDeck()
	dCopy.shuffle()
	if d.toString() == dCopy.toString() {
		t.Errorf("Test failed. Original set: %v,\nShuffled set: %v", d.toString(), dCopy.toString())
	}
}

func Test_deal_success(t *testing.T) {
	hand, remainingDeck := deal(newDeck(), 5)
	if len(hand) != 5 || len(remainingDeck) != 11 {
		t.Errorf("Test failed.")
	}
}
