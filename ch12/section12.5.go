package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	x := 2
	d := reflect.ValueOf(&x).Elem()
	px := d.Addr().Interface().(*int)
	*px = 3
	fmt.Println(x)

	d.Set(reflect.ValueOf(5))
	fmt.Println(x)

	// y := 2
	// b := reflect.ValueOf(y)
	// b.Set(reflect.ValueOf(3)) //panic: reflect: reflect.Value.Set using unaddressable value

	test1()
	test2()
	test3()
}

func test1() {
	x := 1
	rx := reflect.ValueOf(&x).Elem()
	rx.SetInt(3)
	rx.Set(reflect.ValueOf(4))
	// rx.SetString("hello")            // panic: string is not assignable to int
	// rx.Set(reflect.ValueOf("hello")) // panic: string is not assignable to int
}

func test2() {
	var y any
	ry := reflect.ValueOf(&y).Elem()
	// ry.SetInt(3) //panic
	// ry.SetString("hello") //panic: reflect: call of reflect.Value.SetString on interface Value
	ry.Set(reflect.ValueOf(5))
	ry.Set(reflect.ValueOf("hello"))
}

func test3() {
	stdout := reflect.ValueOf(os.Stderr).Elem()
	fmt.Println(stdout.Type())
	pfd := stdout.FieldByName("pfd")
	fd := pfd.FieldByName("Sysfd")
	fmt.Println(fd.Int())
	// fd.SetInt(2) //panic: reflect: reflect.Value.SetInt using value obtained using unexported field
	fmt.Println(fd.CanAddr(), fd.CanSet())
}
