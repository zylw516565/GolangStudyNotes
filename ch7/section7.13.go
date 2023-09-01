package main

import (
	"fmt"
)

func sqlQuoteString(s string) string {
	return s
}

func sqlQuote(x interface{}) string {
	if nil == x {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if s, ok := x.(string); ok {
			return sqlQuoteString(s) // (not shown)
	} else {
			panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func sqlQuote2(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuoteString(x)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func main() {
	fmt.Println("section7.13")
}