package main

import (
	"fmt"
	"bytes"
)

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// return the number of elements
func (s *IntSet) Len() int {
	counter := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word & (1 << uint(j)) != 0 {
				counter++
			}
		}
	}

	return counter
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &= ^(1<<bit)
	}
}

// remove all elements from the set
func (s *IntSet) Clear() {
	for i, _ := range s.words {
		s.words[i] = 0
	}

	// for i := 0; i < len(s.words); i++ {
	// 	*(&s.words[i]) = 0
	// }
}

 // return a copy of the set
func (s *IntSet) Copy() *IntSet {
	d := &IntSet{words: make([]uint64, len(s.words))}
	copy(d.words, s.words)
	return d
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}

	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var d uint8 = 2
	fmt.Printf("%08b\n", d)        // 00000010
	fmt.Printf("%08b\n", ^d)    // 11111101

	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.Len())
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.Len())

	x.Remove(9)
	fmt.Println(x.String())
	fmt.Println(x.Copy())

	x.Clear()
	fmt.Println(x.String())
}