package bitset_test

import (
	"testing"

	"github.com/tomcraven/bitset"
)

func TestCreate(t *testing.T) {
	bitset.Create(10)
}

func createFromStringTest(input string, t *testing.T) {
	b := bitset.CreateFromString(input)
	if b.Size() != uint(len(input)) {
		t.Error("bitset should have size", len(input))
	}

	for i, character := range input {
		shouldBeSet := character == '1'
		shouldBeClear := !shouldBeSet
		if shouldBeSet && !b.Get(uint(i)) {
			t.Error("bitset should have bit at position", i, "set")
		} else if shouldBeClear && b.Get(uint(i)) {
			t.Error("bitset should not have bit at position", i, "set")
		}
	}
}

func TestCreateFromString(t *testing.T) {
	createFromStringTest("00000000", t)
	createFromStringTest("0", t)
	createFromStringTest("1", t)
	createFromStringTest("0110", t)
}
