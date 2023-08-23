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

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) UnionWithMyEdition(t *IntSet) {
	var count int
	if len(s.words) < len(t.words) {
		count = len(s.words)
	} else {
		count = len(t.words)
	}

	for i := 0; i < count; i++ {
		s.words[i] |= t.words[i]
	}

	if len(s.words) < len(t.words) {
		s.words = append(s.words, t.words[count:]...)
	}
}

func (s *IntSet) StringMyEdition() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		if i < 1 {
			fmt.Fprintf(&buf, "%d", word)
			// fmt.Printf("%d", word)
		}else {
			fmt.Fprintf(&buf, " %d", word)
			// fmt.Printf(" %d", word)
		}
	}

	buf.WriteByte('}')
	return buf.String()
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

type Student struct {
	name string
	Age  int
}

func (s *Student) Print() {
	fmt.Println(s.name)
	fmt.Println(s.Age)
}

func main () {
	fmt.Println("IntSet")

	set := &IntSet{[]uint64{7, 4, 5}}
	// set.String()
	fmt.Println(set.StringMyEdition())
	fmt.Println(set.String())
	fmt.Println(^uint(0))

	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	// x.UnionWithMyEdition(&y)
	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(&x) // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)    //{[4398046511618 0 65536]}

	stu := Student{}
	stu.name = "peter"
	stu.Age  = 18
	fmt.Println(stu)
	stu.Print()
}