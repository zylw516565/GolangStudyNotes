package main

import (
	"fmt"
	"time"
	"math"
)

const (
	e = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
	)

const IPv4Len = 4

const (
	a = 4
	b
	c = 2
	d
)

type Weekday int
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint
const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool { return v&FlagUp == FlagUp }
func TurnDown(v *Flags) { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool { return v&(FlagBroadcast|FlagMulticast) != 0 }

const (
	_ = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

var x float32 = math.Pi
var y float64 = math.Pi
var z complex128 = math.Pi

const (
	deadbeef = 0xdeadbeef // untyped int with value 3735928559
	A = uint32(deadbeef) // uint32 with value 3735928559
	B = float32(deadbeef) // float32 with value 3735928576 (rounded up)
	C = float64(deadbeef) // float64 with value 3735928559 (exact)
	// D = int32(deadbeef) // compile error: constant overflows int32
	// E = float64(1e309) // compile error: constant overflows float64
	// F = uint(-1) // compile error: constant underflows uint
	)

func main() {
	fmt.Println(pi)

	fmt.Printf("pi type:%T,\n", pi)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute

	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	fmt.Println(a, b, c, d)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)

	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"

	fmt.Printf("KiB type:%T, MiB type:%T, GiB type:%T, TiB type:%T, PiB type:%T, EiB type:%T\n", KiB, MiB, GiB, TiB, PiB, EiB)
	fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB)

	fmt.Println(YiB/ZiB)
	fmt.Println(x, y, z)

	var f2 float64 = 212
	fmt.Printf("%T,%v\n", (f2-32) * 5, (f2-32) * 5/9)
	fmt.Printf("%T,%v\n", 5/9       , 5/9 * (f2-32))
	fmt.Printf("%T,%v\n", 5.0/9.0   , 5.0/9.0 *(f2-32))

	var f float64 = 3 + 0i
	fmt.Printf("%T,%v\n", 3 + 0i, f)
	f = 2
	fmt.Printf("%T,%v\n", 2, f)
	f = 1e123
	fmt.Printf("%T,%v\n", 1e123, f)
	f = 'a'
	fmt.Printf("%T,%v\n", 'a', f)
	f = '京'
	fmt.Printf("%T,%v\n", '京', f)

	var f3 float64 = float64(3 + 0i)
	_ = f3
	fmt.Println(f3)
	f3 = float64(2)
	fmt.Println(f3)
	f3 = float64(1e123)
	fmt.Println(f3)
	f3 = float64('a')
	fmt.Println(f3)

	vali := 0
	valr := '\000'
	valf := 0.0
	valc := 0i
  fmt.Println(vali, valr, valf, valc)

	fmt.Printf("%T\n", 0) // "int"
	fmt.Printf("%T\n", 0.0) // "float64"
	fmt.Printf("%T\n", 0i) // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)
}


func parseIPv4(s string) (IP string) {
	var p [IPv4Len]byte
	_ = p

	return ""
}
