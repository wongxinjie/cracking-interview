/**
* @File: bitset.go
* @Author: wongxinjie
* @Date: 2019/6/26
*/
package bitset

type BitSet struct {
	bitSet []int
}

func New(capacity int) *BitSet {
	bits := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		bits = append(bits, 0)
	}
	return &BitSet{
		bitSet: bits,
	}
}

func (s *BitSet) Set(pos int) {
	wordNumber := pos >> uint(5)
	bitNumber := pos & 0x1F
	s.bitSet[wordNumber] |= 1 << uint(bitNumber)
}

func (s *BitSet) Get(pos int) int{
	wordNumber := pos >> uint(5)
	bitNumber := pos & 0x1F
	return s.bitSet[wordNumber] & 1 << uint(bitNumber)
}