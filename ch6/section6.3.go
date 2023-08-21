package main

import (
	"fmt"
	"math"
	"sync"
	"image/color"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64{
	fmt.Println("value method: Distance")
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	fmt.Println("pointer method: ScaleBy")
	p.X *= factor
	p.Y *= factor
}

type IntType int

func (p IntType) Distance(q Point) float64{
	fmt.Println("IntType:: Distance")
	return 0
}


type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p ColoredPoint) Distance(q Point) float64{
	fmt.Println("wrapper method: Distance")
	// return p.Point.Distance(q)
	return 0
}

// func (p *ColoredPoint) ScaleBy(factor float64){
// 	fmt.Println("wrapper method: Distance")
// 	p.Point.ScaleBy(factor)
// }

type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

type ColoredPoint3 struct {
	*Point
	Color color.RGBA
	IntType
}

func main () {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"
	fmt.Println("******************************")

	//******************************
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	// fmt.Println(p.Distance(q))  //cannot use q (variable of type ColoredPoint) as Point value in argument to p.Distance
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
	fmt.Println("******************************")

	//******************************
	pPtr := ColoredPoint2{&Point{1, 1}, red}
	qPtr := ColoredPoint2{&Point{5, 4}, blue}
	fmt.Println(pPtr.Distance(*qPtr.Point))
	qPtr.Point = pPtr.Point
	pPtr.ScaleBy(2)
	fmt.Println(*pPtr.Point, *qPtr.Point)
	fmt.Println("******************************")
	// pPtr3 := ColoredPoint3{&Point{1, 1}, red, 0}
	// qPtr3 := ColoredPoint3{&Point{5, 4}, blue, 0}
	// fmt.Println(pPtr3.Distance(*qPtr3.Point))    //ambiguous selector pPtr3.Distance
}

var (
	mu sync.Mutex  // guards mapping
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

var cache = struct {
	sync.Mutex  // guards mapping
	mapping map[string]string
}{mapping: make(map[string]string)}

func Lookup2(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}