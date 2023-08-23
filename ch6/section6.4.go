package main

import (
	"fmt"
	"time"
	"math"
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

type Rocket struct { /* ... */ }
func (r *Rocket) Launch() {
	fmt.Println("Launch")
}



func main () {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))
	var origin Point // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", sqrt(5)

	scaleP := p.ScaleBy
	scaleP(2)
	scaleP(3)
	scaleP(10)
	fmt.Println(p)
	fmt.Println("******************************")
	
	r := new(Rocket)
	time.AfterFunc(2 * time.Second, func() {r.Launch() })
	time.AfterFunc(2 * time.Second, r.Launch)
	// time.Sleep(4 * time.Second)
	fmt.Println("******************************")

	distance := Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p) // "{2 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"
}

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }
type Path []Point
func (path Path) TranslateBy(offset Point, add bool) {
var op func(p, q Point) Point
if add {
op = Point.Add
} else {
op = Point.Sub
}

for i := range path {
// Call either path[i].Add(offset) or path[i].Sub(offset).
path[i] = op(path[i], offset)
}
}