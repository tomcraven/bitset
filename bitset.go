package bitset

type Bitset interface {
	Set(uint)
	Get(uint) bool
	Length() uint
}

type bitset struct {
	bits []uint64
	length uint
}

func Create(size uint) Bitset {
	return &bitset{
		length: size,
	}
}

func (b* bitset) Set(index uint) {
}

func (b* bitset) Get(index uint) bool {
	return false
}

func (b *bitset) Length() uint {
	return b.length
}
