package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 3)

	if sum != 4 {
		t.Error("expected sum to 4")
	}
}
