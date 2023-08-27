package main

import (
	"io"
	"log"
	"fmt"
	"math"
	// "errors"
	"strings"
	"net/http"
)

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

type Var     string
func (v Var)Eval(env Env) float64 {
	return env[v]
}
func (v Var)Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

//***********************************
type literal float64
func (l literal)Eval(_ Env) float64 {
	return float64(l)
}
func (l literal)Check(_ map[Var]bool) error {
	return nil
}

//***********************************
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}
func (u unary)Eval(env Env) float64 {
	switch u.op {
		case '+':
			return +u.x.Eval(env)
		case '-':
			return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}
func (u unary)Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)

//**************My Edition**************
	// switch u.op {
	// case '+':
	// 	fallthrough
	// case '-':
	// 	return u.x.Check(vars)
	// }

	// return fmt.Errorf("unsupported unary operator")
}

//***********************************
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}
func (b binary)Eval(env Env) float64 {
	switch b.op {
		case '+':
			return b.x.Eval(env) + b.y.Eval(env)
		case '-':
			return b.x.Eval(env) - b.y.Eval(env)
		case '*':
			return b.x.Eval(env) * b.y.Eval(env)
		case '/':
			return b.x.Eval(env) / b.y.Eval(env)
	}

	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}
func (b binary)Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)

//**************My Edition**************
	// switch b.op {
	// case '+':
	// 	fallthrough
	// case '-':
	// 	fallthrough
	// case '*':
	// 	fallthrough
	// case '/':
	// 	return errors.Join(b.x.Check(vars), b.y.Check(vars))
	// }

	// return fmt.Errorf("unsupported binary operator")
}

//***********************************
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}
func (c call)Eval(env Env) float64 {
	switch c.fn {
		case "pow":
			return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
		case "sin":
			return math.Sin(c.args[0].Eval(env))
		case "sqrt":
			return math.Sqrt(c.args[0].Eval(env))
	}

	panic(fmt.Sprintf("unsupported call operator: %q", c.fn))
}
func (c call)Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}

	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
		c.fn, len(c.args), arity)
	}

	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
	//**************My Edition**************
	// var err error
	// switch c.fn {
	// case "sin":
	// 	fallthrough
	// case "sqrt":
	// 	if len(c.args) != 1 {
	// 		err = fmt.Errorf("len(args) != 1")
	// 	}

	// case "pow":
	// 	if len(c.args) != 2 {
	// 		err = fmt.Errorf("len(args) != 2")
	// 	}
	// default:
	// 	return fmt.Errorf("unknown function %q", c.fn)
	// }

	// for _, arg := range c.args {
	// 	err = errors.Join(err, arg.Check(vars))
	// }

	// return err
}

//***********************************
type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

type Env map[Var]float64



//***********************************
//***********************************
const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // x, y axis range (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
)

var sin30, cos30 = 0.5, math.Sqrt(3.0 / 4.0) // sin(30°), cos(30°)

func corner(f func(x, y float64) float64, i, j int) (float64, float64) {
	// find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y) // compute surface height z

	// project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func surface(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(f, i+1, j)
			bx, by := corner(f, i, j)
			cx, cy := corner(f, i, j+1)
			dx, dy := corner(f, i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func parseAndCheck(s string) (Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance from (0,0)
		return expr.Eval(Env{"x": x, "y": y, "r": r})
	})
}

func main() {
	http.HandleFunc("/plot", plot)
	log.Fatal(http.ListenAndServe("172.26.88.88:8000", nil))
}

//http://127.0.0.1:8000/plot?expr=sin(-x)*pow(1.5, -r)
//http://172.26.88.88:8000/plot?expr=sin(-x)*pow(1.5,-r)
//http://172.26.88.88:8000/plot?expr=pow(2,sin(y))*pow(2,sin(x))/12
//http://172.26.88.88:8000/plot?expr=sin(x*y/10)/10