package bitset

type Bitset interface {
	Set(int)
	Has(int) bool
}

type bitset struct {
	set bool
}

func Create(size int) Bitset {
	return &bitset{}
}


func (b* bitset) Set(index int) {
	b.set = true
}

func (b* bitset) Has(index int) bool {
	return b.set
}
