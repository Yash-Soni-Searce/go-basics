package main

import (
	"os"
	"testing"
)

func Testdeck(t *testing.T) {
	d := newdeckfromfile("my-cards")
	if len(d) != 52 {
		t.Errorf("Need 52 but got %v", len(d))
	}
}

func TestSaveToFileAndnewdeckfromfile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newdeck()
	deck.savefile("_decktesting", deck.tobytestring())
	loadeddeck := newdeckfromfile("_decktesting")

	if len(loadeddeck) != 52 {
		t.Errorf("Needed 52 cards but got %v", len(loadeddeck))
	}
	os.Remove("_decktesting")

}
