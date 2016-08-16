package bitset

import "fmt"

type slice struct {
	other      Bitset
	begin, end uint
}

func (s *slice) Set(indexes ...uint) {
	for i := uint(0); i < uint(len(indexes)); i++ {
		index := indexes[i]
		s.other.Set(s.begin + index)
	}
}

func (s *slice) SetTo(index uint, val bool) {
	s.other.SetTo(s.begin+index, val)
}

func (s *slice) SetAll() {
	for i := s.begin; i < s.end; i++ {
		s.other.Set(i)
	}
}

func (s *slice) Clear(indexes ...uint) {
	for i := uint(0); i < uint(len(indexes)); i++ {
		index := indexes[i]
		s.other.Clear(s.begin + index)
	}
}

func (s *slice) ClearAll() {
	for i := s.begin; i < s.end; i++ {
		s.other.Clear(i)
	}
}

func (s *slice) Invert() {
	for i := s.begin; i < s.end; i++ {
		s.other.SetTo(i, !s.other.Get(i))
	}
}

func (s *slice) Get(index uint) bool {
	return s.other.Get(s.begin + index)
}

func (s *slice) Size() uint {
	return s.end - s.begin
}

func (s *slice) Clone() Bitset {
	return &slice{
		begin: s.begin,
		end:   s.end,
		other: s.other.Clone(),
	}
}

func (s *slice) Slice(begin, end uint) Bitset {
	return &slice{
		begin: begin,
		end:   end,
		other: s,
	}
}

func (s *slice) Equals(other Bitset) bool {
	return other.equalsSlice(s)
}

func (s *slice) equalsBitset(other *bitset) bool {
	return equalsBitsetSlice(other, s)
}

func (s *slice) equalsSlice(other *slice) bool {
	return equalsSliceSlice(s, other)
}

func (s *slice) Output() {
	for i := uint(0); i < s.Size(); i++ {
		if s.Get(i) {
			fmt.Printf("1")
		} else {
			fmt.Printf("0")
		}
	}
	fmt.Printf("\n")
}

func (s *slice) BuildUint8(output *uint8) bool {
	if s.Size() < 8 {
		return false
	}

	for i := uint(0); i < 8; i++ {
		if s.Get(i) {
			(*output) |= 1 << i
		} else {
			(*output) &= ^(1 << i)
		}
	}
	return true
}
