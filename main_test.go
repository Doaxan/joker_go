package main

import (
	"testing"
)

// TestRandomJoke calls randomJoke(),
// checking for a valid return value.
func TestRandomJoke(t *testing.T) {
	value := randomJoke()
	if len(value) < 1 {
		t.Fatalf("randomJoke() = %v, want joke, error", value)
	}
}

// TestRandomJokeCategory calls randomJokeCategory() with a category,
// checking for a valid return value.
func TestRandomJokeCategory(t *testing.T) {
	value := randomJokeCategory("animal")
	if len(value) < 1 {
		t.Fatalf("randomJokeCategory() = %v, want joke, error", value)
	}
}

