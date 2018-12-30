package main

import (
	"os"
	"testing"
)

func TestNewDesk(t *testing.T) {
	d := newDesk()

	if len(d) != 16 {
		t.Errorf("Expected 16 but got %v", len(d))
	}

	if d[0] != "ace of H" {
		t.Errorf("Expected ace of H, get %v", d[0])
	}

	if d[len(d)-1] != "3 of C" {
		t.Errorf("Expected 3 of C, got %v", d[len(d)-1])
	}
}

func TestDeskFromFile(t *testing.T) {
	filename := "_desktesting.txt"
	os.Remove(filename)
	// Add new desk to file
	desk := newDesk()
	desk.saveToFile(filename)
	// Load desk from file
	loadedDesk := newDeskFromFile(filename)

	// Test cases
	if len(loadedDesk) != 16 {
		t.Errorf("Expected 16 but got %v", len(loadedDesk))
	}

	os.Remove(filename)
}
