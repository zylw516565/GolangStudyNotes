package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64{
	fmt.Println("function call")
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point)Distance(q Point) float64{
	fmt.Println("method call")
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point
// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}

	return sum
}

type IntType int

func (i *IntType) printInt() {
	fmt.Println(*i)
}

// type P *int
// func (P) f(){}    //invalid receiver type P

func main () {
	p := Point{1, 2}
	q := Point{3, 4}

	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	fmt.Println((&p).Distance(q))
	pPtr := &p
	fmt.Println(pPtr.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"

	var a IntType = 10
	a.printInt()
}