package bitset

const bitsPerUint64 = 64

type Bitset interface {
	Set(...uint)
	SetTo(uint, bool)
	SetAll()
	Clear(...uint)
	ClearAll()
	Invert()
	Get(uint) bool
	Size() uint
	Clone() Bitset
	Slice(uint, uint) Bitset
	Output()
	BuildUint8(*uint8) bool

	Equals(Bitset) bool
	equalsBitset(*bitset) bool
	equalsSlice(*slice) bool
}

func bitArraySize(numBits uint) uint {
	return (numBits / bitsPerUint64) + 1
}

// Create takes a size and returns a bitset with that size
func Create(size uint) Bitset {
	b := bitset{}
	b.init(size)
	return &b
}

// CreateFromString takes a string of 1s and 0s and returns a bitset
// that matches that
func CreateFromString(str string) Bitset {
	b := Create(uint(len(str)))
	for i, char := range str {
		b.SetTo(uint(i), char == '1')
	}
	return b
}

func equalsBitsetBitset(a *bitset, b *bitset) bool {
	if a.Size() != b.Size() {
		return false
	}

	for i := 0; i < len(a.bits); i++ {
		if a.bits[0] != b.bits[0] {
			return false
		}
	}

	return true
}

// TODO - can we do better?
func naiveEquality(a Bitset, b Bitset) bool {
	if a.Size() != b.Size() {
		return false
	}

	for i := uint(0); i < a.Size(); i++ {
		if a.Get(i) != b.Get(i) {
			return false
		}
	}

	return true
}

func equalsBitsetSlice(a *bitset, b *slice) bool {
	return naiveEquality(a, b)
}

func equalsSliceSlice(a *slice, b *slice) bool {
	return naiveEquality(a, b)
}
