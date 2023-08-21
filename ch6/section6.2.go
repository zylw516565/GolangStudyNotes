package main

import (
	"fmt"
	"math"
	"net/url"
)

type Point struct {
	X, Y float64
}

func (p Point)Distance(q Point) float64{
	fmt.Println("value method: Distance")
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	fmt.Println("pointer method: ScaleBy")
	p.X *= factor
	p.Y *= factor
}

// can not overload method
// func (p Point) ScaleBy(factor float64) {
// 	fmt.Println("value method: ScaleBy")
// 	p.X *= factor
// 	p.Y *= factor
// }

// type P *int
// func (P) f(){}    //invalid receiver type P

func main () {

	//call format 1
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r) // "{2, 4}"

	//call format 2
	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	(&*pptr).ScaleBy(2)
	fmt.Println(p) // "{2, 4}"

	//call format 3
	(&p).ScaleBy(2)
	fmt.Println(p)

	// Point{1, 2}.ScaleBy(2)    //cannot call pointer method ScaleBy on Point
	p.ScaleBy(2)
	fmt.Println(p)

	q := Point{3, 4}
	fmt.Println(Point{3, 4}.Distance(q))
	fmt.Println(pptr.Distance(q))
	fmt.Println((*pptr).Distance(q))

	testUrlValues()
}

type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}

	return list.Value + list.Tail.Sum()
}

func testUrlValues() {
	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q")) // ""
	fmt.Println(m.Get("item")) // "1" (first value)
	fmt.Println(m["item"]) // "[1 2]" (direct map access)
	m = nil
	fmt.Println(m.Get("item")) // ""
	fmt.Println(url.Values(nil).Get("item")) // ""
	// m.Add("item", "3") // panic: assignment to entry in nil map
}