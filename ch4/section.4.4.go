package main

import (
	"fmt"
	"time"
	"GolangStudyNotes/ch4/somepackage"
)

type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

type Employee2 struct {
	ID int
	Name, Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

var dilbert Employee

type S struct {
	ID   int
	Name string
	s    *S
	// s    S   //invalid recursive type: S refers to itself
}

type Point struct{ X, Y int }

type address struct {
	hostname string
	port     int
}

func main() {
	dilbert.Salary -= 5000

	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"

	(*employeeOfTheMonth).Position = " (proactive team player aaa)"

	EmployeeByID(0).Salary = 0
	// EmployeeByIDV2(0).Salary = 0

	fmt.Printf("%v\n", dilbert)

	p := Point{1,13}
	p2 := Point{Y:1, X:22}
	fmt.Printf("%v\n", p)
	fmt.Println(p2)

	// some := somepackage.T{a:2, b:3}     //unknown field a in struct literal of type somepackage.T
	// fmt.Println(some)
	// some2 := somepackage.T{2, 3}        //implicit assignment to unexported field a in struct literal of type somepackage.T
	// fmt.Println(some2)

	p3 := Point{3, 3}
	factor := 2
	fmt.Println(Scale(p3, factor), Scale2(p3, factor), p3)
	fmt.Println(ScaleWithPtr(&p3, factor), p3)

	p4 := &Point{1, 2}
	fmt.Println(p4)
	p5 := new(Point)
	*p5 = Point{3, 4}
	fmt.Println(p5)

	r := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(r.X == q.X && r.Y == q.Y) // "false"
	fmt.Println(r == q) // "false"

	addrs := make(map[address]int)
	addrs[address{"baidu.com", 80}]++
	fmt.Println(addrs)

	StructTest1()
	StructTest2()
	StructTest3()
	StructTest4()
}

func EmployeeByID(id int) *Employee {
	dilbert := &Employee{}

	return dilbert
}

func EmployeeByIDV2(id int) Employee {
	dilbert := Employee{}

	return dilbert
}

func Scale(in Point, factor int) (out Point) {
	return Point{X:in.X*factor, Y:in.Y*factor }
}

func Scale2(in Point, factor int) Point {
	return Point{in.X*factor, in.Y*factor }
}

func ScaleWithPtr(in *Point, factor int) *Point {
	in.X *= factor
	in.Y *= factor
	return in
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

func StructTest1() {
	type Circle struct {
		X, Y, Radius int
	}
	
	type Wheel struct {
		X, Y, Radius, Spokes int
	}

	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	fmt.Println(w)
}

func StructTest2() {
	type Point struct {
		X, Y int
	}

	type Circle struct {
		Center Point
		Radius int
	}

	type Wheel struct {
		Circle Circle
		Spokes int
	}

	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
	fmt.Println(w)
}

func StructTest3() {
	type Point struct {
		X, Y int
	}

	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}

	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	w.Circle.Point.X = 9
	w.Circle.Point.Y = 9
	w.Circle.Radius = 6
	w.Spokes = 21
	fmt.Println(w)

	// w2 := Wheel{X:1, Y:1, Radius:1, Spokes: 1} //error?
	// fmt.Println(w2)

	// w3 := Wheel{}
	w3 := Wheel{Circle{Point{2, 2}, 2}, 2}
	fmt.Println(w3)

	w4 := Wheel{
		Circle: Circle {
			Point: Point{X:2, Y:2},
			Radius:2,
		},
		Spokes:2,
	}
	fmt.Println(w4)
	fmt.Printf("%v\n", w4)
	fmt.Printf("%#v\n", w4)
}

func StructTest4() {
	var w somepackage.White
	w.X = 5
	w.Y = 5
	w.Radius = 5
	w.Spokes = 5

	// w.Circle.Point.X = 9
	// w.Circle.Point.Y = 9
	// w.Circle.Radius = 6
	// w.Spokes = 21
	fmt.Println(w)
}