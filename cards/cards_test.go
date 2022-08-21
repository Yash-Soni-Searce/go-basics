package main

import "testing"

func Testdeck(t *testing.T) {
	d := newdeckfromfile("my-cards")
	if len(d) != 52 {
		t.Errorf("Need 52 but got %v", len(d))
	}
}
