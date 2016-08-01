package bitset_test

import "testing"
import "github.com/tomcraven/bitset"

func TestCreate(t *testing.T) {
	bitset.Create(10)
}

func TestSetGet(t *testing.T) {
	bitset := bitset.Create(10)
	if bitset.Get(0) {
		t.Error("bitset should not have the bit at position 0 set")
	}
	if bitset.Get(5) {
		t.Error("bitset should not have the bit at position 0 set")
	}

	bitset.Set(0)
	bitset.Set(5)
	if !bitset.Get(0) {
		t.Error("bitset should have the bit at position 0 set")
	}
	if !bitset.Get(5) {
		t.Error("bitset should have the bit at position 0 set")
	}
}

func sizeTest(size uint, t *testing.T) {
	bitset := bitset.Create(size)
	if bitset.Size() != size {
		t.Error("bitset should have size", size)
	}
}

func TestSize(t *testing.T) {
	sizeTest(0, t)
	sizeTest(10, t)
}

func TestClear(t *testing.T) {
	bitset := bitset.Create(10)

	bitset.Set(0)
	bitset.Set(5)
	bitset.Clear(0)

	if bitset.Get(0) {
		t.Error("bitset should not have the bit at position 0 set")
	}
	if !bitset.Get(5) {
		t.Error("bitset should have the bit at position 0 set")
	}
}
