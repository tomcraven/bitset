package bitset_test

import "testing"
import "github.com/tomcraven/bitset"

func TestSet(t *testing.T) {
	bitset := bitset.Create(10)
	bitset.Set(0)
}

func TestHas(t *testing.T) {
	bitset := bitset.Create(10)
	if bitset.Has(0) {
		t.Error("bitset should not have the bit at position 0 set")
	}

	bitset.Set(0)
	if !bitset.Has(0) {
		t.Error("bitset should have the bit at position 0 set")
	}
}
