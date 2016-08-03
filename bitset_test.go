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

func TestCreateCopy(t *testing.T) {
	bitset := bitset.Create(10)
	bitset.Set(0)
	bitset.Set(3)
	bitset.Set(8)

	copy := bitset.CreateCopy()
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

func TestEquals(t *testing.T) {
	a := bitset.Create(10)
	b := bitset.Create(10)
	c := bitset.Create(20)

	if a.Equals(c) {
		t.Error("a should not equal c")
	}

	for _, v := range []uint{0, 5, 6, 7} {
		a.Set(v)
		b.Set(v)
	}

	if !a.Equals(b) {
		t.Error("a should equal b")
	}

	b.ClearAll()

	if a.Equals(b) {
		t.Error("a should now not equal b")
	}
}
