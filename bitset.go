package bitset

import (
	"fmt"
	"math"
)

const bitsPerUint64 = 64 // sizeof(uint64) * 8

type Bitset interface {
	Set(uint)
	SetTo(uint, bool)
	SetAll()

	Clear(uint)
	ClearAll()

	Get(uint) bool

	Size() uint
	CreateCopy() Bitset

	Equals(Bitset) bool
	equalsBitset(*bitset) bool

	Output()
}

type bitset struct {
	bits []uint64
	size uint
}

func bitArraySize(numBits uint) uint {
	return (numBits / bitsPerUint64) + 1
}

func Create(size uint) Bitset {
	b := bitset{}
	b.init(size)
	return &b
}

func (b *bitset) init(size uint) {
	b.size = size
	b.bits = make([]uint64, bitArraySize(size))
}

func (b *bitset) Set(index uint) {
	elementIndex := index / bitsPerUint64
	bitIndex := index % bitsPerUint64
	b.bits[elementIndex] |= 1 << bitIndex
}

func (b *bitset) SetTo(index uint, value bool) {
	if value {
		b.Set(index)
	} else {
		b.Clear(index)
	}
}

func (b *bitset) setImpl(index uint) {
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

func (b *bitset) CreateCopy() Bitset {
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

func (b *bitset) equalsBitset(other *bitset) bool {
	if b.Size() != other.Size() {
		return false
	}

	for i, v := range b.bits {
		if v != other.bits[i] {
			return false
		}
	}

	return true
}

func (b *bitset) Equals(other Bitset) bool {
	return other.equalsBitset(b)
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
