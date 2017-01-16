package bitset_test

import "testing"

import "github.com/tomcraven/bitset"

func shouldBeSet(t *testing.T, b bitset.Bitset, indexes ...uint) {
	for _, v := range indexes {
		if !b.Get(v) {
			t.Error("bitset should have bit set at position", v)
		}
	}
}

func shouldBeClear(t *testing.T, b bitset.Bitset, indexes ...uint) {
	for _, v := range indexes {
		if b.Get(v) {
			t.Error("bitset should not have bit set at position", v)
		}
	}
}

func TestSliceGetSetClear(t *testing.T) {
	b := bitset.Create(10)
	b.Set(0, 1, 2, 3, 4)

	slice := b.Slice(0, 5)
	if slice.Size() != 5 {
		t.Error("sliced bitset should have size", 5)
	}
	shouldBeSet(t, slice, 0, 1, 2, 3, 4)

	// Should set and clear bits in original bitset
	slice.Clear(0, 1, 2, 3, 4)
	slice.Set(5, 6, 7, 8, 9)

	for i := uint(0); i < 5; i++ {
		if slice.Get(i) || b.Get(i) {
			t.Error("original and sliced bitset should not have bit set at index", i)
		}
	}

	for i := uint(5); i < 10; i++ {
		if !(slice.Get(i) && b.Get(i)) {
			t.Error("original and sliced btiset should have bit set at index", i)
		}
	}
}

func TestSliceSetTo(t *testing.T) {
	b := bitset.Create(10)
	b.Set(5, 6, 7, 8, 9)

	slice := b.Slice(0, 5)
	slice.SetTo(0, true)
	slice.SetTo(5, false)

	shouldBeSet(t, b, 0)
	shouldBeClear(t, b, 5)
}

func TestSliceSetAll(t *testing.T) {
	b := bitset.Create(10)
	slice := b.Slice(2, 7)
	slice.SetAll()

	shouldBeSet(t, b, 2, 3, 4, 5, 6)
	shouldBeClear(t, b, 0, 1, 7, 8, 9)
}

func TestSliceClearAll(t *testing.T) {
	b := bitset.Create(10)
	b.SetAll()
	slice := b.Slice(2, 7)
	slice.ClearAll()

	shouldBeSet(t, b, 0, 1, 7, 8, 9)
	shouldBeClear(t, b, 2, 3, 4, 5, 6)
}

func TestSliceInvert(t *testing.T) {
	b := bitset.Create(10)
	slice := b.Slice(2, 7)
	slice.Invert()

	shouldBeSet(t, b, 2, 3, 4, 5, 6)
	shouldBeClear(t, b, 0, 1, 7, 8, 9)
}

func TestSliceClone(t *testing.T) {
	b := bitset.Create(10)
	slice := b.Slice(2, 7)
	slice.Set(0, 3, 4)

	sliceClone := slice.Clone()
	sliceClone.Invert()

	if sliceClone.Size() != slice.Size() {
		t.Error("slice and clone should have the same size")
	}

	for i := uint(0); i < slice.Size(); i++ {
		if slice.Get(i) == sliceClone.Get(i) {
			t.Error("slice clone was inverted, it should be opposite to original")
		}
	}
}

func TestSliceSlice(t *testing.T) {
	b := bitset.Create(10)
	slice := b.Slice(2, 7)
	sliceSlice := slice.Slice(1, 5)
	sliceSlice.SetAll()

	// sliceSlice == 1111
	shouldBeSet(t, sliceSlice, 0, 1, 2, 3)

	// slice == 0111100
	shouldBeClear(t, slice, 0, 5, 6)
	shouldBeSet(t, slice, 1, 2, 3, 4)

	// b = 0001111000
	shouldBeClear(t, b, 0, 1, 2, 7, 8, 9)
	shouldBeSet(t, b, 3, 4, 5, 6)
}

func TestSliceEqualsBitset(t *testing.T) {
	b := bitset.Create(10)
	slice := b.Slice(0, 10)
	newBitset := bitset.Create(10)

	b.Set(0, 2)
	slice.Set(1, 3)
	if !slice.Equals(b) {
		t.Error("bitset and slice should be equal")
	}

	newBitset.Set(0, 1, 2, 3)
	if !slice.Equals(newBitset) {
		t.Error("slice and new bitset should be equal")
	}

	newBitset.ClearAll()
	if slice.Equals(newBitset) {
		t.Error("slice and new bitset should not be equal")
	}
}

func TestSliceEqualsSlice(t *testing.T) {
	a := bitset.Create(10).Slice(0, 10)
	b := bitset.Create(10).Slice(0, 10)
	c := b.Slice(0, 2)

	if b.Equals(c) {
		t.Error("slices with different sizes should not be equal")
	}

	if !a.Equals(b) {
		t.Error("slices a and b should be equal")
	}

	a.Set(7, 8, 9)
	if a.Equals(b) {
		t.Error("slices a and b should not be equal")
	}
}

func TestSliceBuildUint8(t *testing.T) {
	// Bitset too small
	buildUint8Test(true, func() bitset.Bitset {
		b := bitset.Create(1)
		return b.Slice(0, 1)
	}, 0, t)

	// Bitset big enough but blank
	buildUint8Test(false, func() bitset.Bitset {
		b := bitset.Create(8)
		b.ClearAll()
		return b.Slice(0, 8)
	}, 0, t)

	// Bitset with all some bits set
	buildUint8Test(false, func() bitset.Bitset {
		b := bitset.Create(8)
		b.SetAll()
		return b.Slice(0, 8)
	}, 255, t)

	// Bitset with some bits set 00000000 01100001 == 97 == 'a'
	buildUint8Test(false, func() bitset.Bitset {
		b := bitset.Create(16)
		b.ClearAll()
		b.Set(8)
		b.Set(13)
		b.Set(14)
		return b.Slice(8, 16)
	}, 97, t)
}

func TestSliceSetBeyondSize(t *testing.T) {
	b := bitset.Create(63)
	slice := b.Slice(0, b.Size())
	slice.Set(64)
}

func TestSliceClearBeyondSize(t *testing.T) {
	b := bitset.Create(63)
	slice := b.Slice(0, b.Size())
	slice.Clear(64)
}

func TestSliceGetBeyondSize(t *testing.T) {
	b := bitset.Create(63)
	b.SetAll()
	slice := b.Slice(0, 10)
	if slice.Get(slice.Size()) {
		t.Error("getting beyond the size of the bitset should return false")
	}
}
