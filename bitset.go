package bitset

const bitsPerUint64 = 64 // sizeof(uint64) * 8

type Bitset interface {
	Set(uint)
	Get(uint) bool
	Clear(uint)
	Size() uint
}

type bitset struct {
	bits []uint64
	size uint
}

func bitArraySize(numBits uint) uint {
	return (numBits / bitsPerUint64) + 1
}

func Create(size uint) Bitset {
	return &bitset{
		size: size,
		bits: make([]uint64, bitArraySize(size)),
	}
}

func (b *bitset) Set(index uint) {
	elementIndex := index / bitsPerUint64
	bitIndex := index % bitsPerUint64
	b.bits[elementIndex] |= 1 << bitIndex
}

func (b *bitset) Get(index uint) bool {
	elementIndex := index / bitsPerUint64
	bitIndex := index % bitsPerUint64
	return ((b.bits[elementIndex] >> bitIndex) & 1) == 1
}

func (b *bitset) Clear(index uint) {
	elementIndex := index / bitsPerUint64
	bitIndex := index % bitsPerUint64
	b.bits[elementIndex] &= ^(1 << bitIndex)
}

func (b *bitset) Size() uint {
	return b.size
}
