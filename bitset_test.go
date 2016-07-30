package bitset_test

import "testing"
import "github.com/tomcraven/bitset"

func TestCreate(t *testing.T) {
	
}

func TestSetGet(t *testing.T) {
	bitset := bitset.Create(10)
	if bitset.Get(0) {
		t.Error("bitset should not have the bit at position 0 set")
	}

	bitset.Set(0)
	if !bitset.Get(0) {
		t.Error("bitset should have the bit at position 0 set")
	}
	
	if bitset.Get(5) {
		t.Error("bitset should not have the bit at position 0 set")
	}
}

func lengthTest(length uint, t *testing.T) {
	bitset := bitset.Create(length)
	if bitset.Length() != length {
		t.Error("bitset should have length", length)
	}
}

func TestLength(t *testing.T) {
	lengthTest(0, t)
	lengthTest(10, t)
}
