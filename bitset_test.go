package bitset_test

import "testing"

import "github.com/tomcraven/bitset"

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

func TestSetTo(t *testing.T) {
	bitset := bitset.Create(10)

	bitset.SetTo(0, false)
	bitset.SetTo(5, true)

	if bitset.Get(0) {
		t.Error("bitset should not have the bit at position 0 set")
	}
	if !bitset.Get(5) {
		t.Error("bitset should have the bit at position 0 set")
	}
}

func TestClone(t *testing.T) {
	bitset := bitset.Create(10)
	bitset.Set(0)
	bitset.Set(3)
	bitset.Set(8)

	copy := bitset.Clone()
	for _, i := range []uint{0, 3, 8} {
		if !copy.Get(i) {
			t.Error("bitset should not have the bit at position", i, "set")
		}
	}

	for _, i := range []uint{1, 2, 4, 5, 6, 7, 9} {
		if copy.Get(i) {
			t.Error("bitset should not have the bit at position", i, "set")
		}
	}
}

func TestSetAll(t *testing.T) {
	bitset := bitset.Create(64)

	for i := uint(0); i < bitset.Size(); i++ {
		if bitset.Get(i) {
			t.Error("bitset should not have the bit at position", i, "set")
		}
	}

	bitset.SetAll()

	for i := uint(0); i < bitset.Size(); i++ {
		if !bitset.Get(i) {
			t.Error("bitset should have the bit at position", i, "set")
		}
	}
}

func TestClearAll(t *testing.T) {
	bitset := bitset.Create(10)
	bitset.SetAll()

	for i := uint(0); i < bitset.Size(); i++ {
		if !bitset.Get(i) {
			t.Error("bitset should have the bit at position", i, "set")
		}
	}

	bitset.ClearAll()

	for i := uint(0); i < bitset.Size(); i++ {
		if bitset.Get(i) {
			t.Error("bitset should not have the bit at position", i, "set")
		}
	}
}

func TestInvert(t *testing.T) {
	b := bitset.Create(10)
	b.Invert()

	for i := uint(0); i < b.Size(); i++ {
		if !b.Get(i) {
			t.Error("bitset should have the bit at position", i, "set")
		}
	}
}

func TestEqualsBitset(t *testing.T) {
	a := bitset.Create(10)
	b := bitset.Create(10)
	c := bitset.Create(20)

	if a.Equals(c) {
		t.Error("a should not equal c")
	}

	a.Set(0, 5, 6, 7)
	b.Set(0, 5, 6, 7)

	if !a.Equals(b) {
		t.Error("a should equal b")
	}

	b.ClearAll()

	if a.Equals(b) {
		t.Error("a should now not equal b")
	}
}

func TestEqualsSlice(t *testing.T) {
	b := bitset.Create(10)
	b.Set(0, 2, 4, 6, 8)
	slice := b.Slice(0, 10)

	if !b.Equals(slice) {
		t.Error("bitset and slice should be equal")
	}

	// Changing these change for both because it's a slice
	// b == slice == 1000000001
	b.Set(9)
	slice.Set(0)
	if !b.Equals(slice) {
		t.Error("bitset and slice should be equal")
	}

	// newBitset == 1000000000
	newBitset := bitset.Create(10)
	newBitset.Set(0)
	if newBitset.Equals(slice) {
		t.Error("newBitset and the old slice should not be equal")
	}
}

func buildUint8Test(shouldFail bool, f func() bitset.Bitset, expected uint8, t *testing.T) {
	b := f()
	var input uint8
	if b.BuildUint8(&input) && shouldFail {
		t.Error("should fail to build")
	}

	if input != expected {
		t.Error(input, "does not equal the expected result", expected)
	}
}

func TestBuildUint8(t *testing.T) {
	// Bitset too small
	buildUint8Test(true, func() bitset.Bitset {
		b := bitset.Create(1)
		return b
	}, 0, t)

	// Bitset big enough but blank
	buildUint8Test(false, func() bitset.Bitset {
		b := bitset.Create(8)
		b.ClearAll()
		return b
	}, 0, t)

	// Bitset with all some bits set
	buildUint8Test(false, func() bitset.Bitset {
		b := bitset.Create(8)
		b.SetAll()
		return b
	}, 255, t)

	// Bitset with some bits set 01100001 == 97 == 'a'
	buildUint8Test(false, func() bitset.Bitset {
		b := bitset.Create(8)
		b.ClearAll()
		b.Set(0)
		b.Set(5)
		b.Set(6)
		return b
	}, 97, t)
}

func TestSlice(t *testing.T) {
	b := bitset.Create(10)
	b.Set(0)
	b.Set(1)
	b.Set(2)

	slice := b.Slice(0, 3)
	if slice.Size() != 3 {
		t.Error("sliced bitset should have size", 3)
	}
	for i := uint(0); i < slice.Size(); i++ {
		if !slice.Get(i) {
			t.Error("sliced bitset should have bit set at index", i)
		}
	}
}

func TestSetVarArgs(t *testing.T) {
	b := bitset.Create(10)
	b.Set(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	for i := uint(0); i < b.Size(); i++ {
		if !b.Get(i) {
			t.Error("bitset should have bit set at index", i)
		}
	}
}

func TestClearVarArgs(t *testing.T) {
	b := bitset.Create(10)
	b.SetAll()
	b.Clear(0, 1, 2, 3, 4)

	for i := uint(0); i < 5; i++ {
		if b.Get(i) {
			t.Error("bitset should not have bitset at index", i)
		}
	}
}

func TestSetBeyondSize(t *testing.T) {
	b := bitset.Create(63)
	b.Set(64)
}

func TestClearBeyondSize(t *testing.T) {
	b := bitset.Create(63)
	b.Clear(64)
}

func TestGetBeyondSize(t *testing.T) {
	b := bitset.Create(63)
	if b.Get(64) {
		t.Error("getting beyond the size of the bitset should return false")
	}
}
