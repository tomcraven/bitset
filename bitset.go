package bitset

import (
	"fmt"
	"math"
)

type bitset struct {
	bits []uint64
	size uint
}

func (b *bitset) init(size uint) {
	b.size = size
	b.bits = make([]uint64, bitArraySize(size))
}

func (b *bitset) Set(indexes ...uint) {
	for i := uint(0); i < uint(len(indexes)); i++ {
		index := indexes[i]
		if index >= b.Size() {
			continue
		}

		elementIndex := index / bitsPerUint64
		bitIndex := index % bitsPerUint64
		b.bits[elementIndex] |= 1 << bitIndex
	}
}

func (b *bitset) SetTo(index uint, value bool) {
	if value {
		b.Set(index)
	} else {
		b.Clear(index)
	}
}

func (b *bitset) Get(index uint) bool {
	if index > b.Size() {
		return false
	}

	elementIndex := index / bitsPerUint64
	bitIndex := index % bitsPerUint64
	return ((b.bits[elementIndex] >> bitIndex) & 1) == 1
}

func (b *bitset) Clear(indexes ...uint) {
	for i := uint(0); i < uint(len(indexes)); i++ {
		index := indexes[i]
		if index > b.Size() {
			continue
		}

		elementIndex := index / bitsPerUint64
		bitIndex := index % bitsPerUint64
		b.bits[elementIndex] &= ^(1 << bitIndex)
	}
}

func (b *bitset) Size() uint {
	return b.size
}

func (b *bitset) Clone() Bitset {
	new := bitset{}
	new.init(b.size)
	copy(new.bits, b.bits)
	return &new
}

func (b *bitset) SetAll() {
	for i := range b.bits {
		b.bits[i] = math.MaxUint64
	}
}

func (b *bitset) ClearAll() {
	for i := range b.bits {
		b.bits[i] = 0
	}
}

func (b *bitset) Invert() {
	for i := range b.bits {
		b.bits[i] = ^b.bits[i]
	}
}

func (b *bitset) Equals(other Bitset) bool {
	return other.equalsBitset(b)
}

func (b *bitset) equalsBitset(other *bitset) bool {
	return equalsBitsetBitset(b, other)
}

func (b *bitset) equalsSlice(other *slice) bool {
	return equalsBitsetSlice(b, other)
}

func (b *bitset) Output() {
	for i := uint(0); i < b.Size(); i++ {
		if b.Get(i) {
			fmt.Printf("1")
		} else {
			fmt.Printf("0")
		}
	}
	fmt.Printf("\n")
}

func (b *bitset) BuildUint8(output *uint8) bool {
	if b.Size() < 8 {
		return false
	}

	for i := uint(0); i < 8; i++ {
		if b.Get(i) {
			(*output) |= 1 << i
		} else {
			(*output) &= ^(1 << i)
		}
	}
	return true
}

func (b *bitset) Slice(begin, end uint) Bitset {
	return &slice{
		begin: begin,
		end:   end,
		other: b,
	}
}
