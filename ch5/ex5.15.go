package main

import (
	"os"
	"fmt"
)

func main () {
	// max, err := max()
	max, err := max(1,2,3,4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(max)

	fmt.Println(min(1,2,3,4))

	fmt.Println(max1(7,8,9))
	fmt.Println(min1(5,6,7))	
}

func max(vals ...int) (int, error) {
	if len(vals) <= 0 {
		return -1, fmt.Errorf("No input parameter")
	}

	max := vals[0]
	for _, val := range vals[1:] {
		if val > max {
			max = val
		}
	}

	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) <= 0 {
		return -1, fmt.Errorf("No input parameter")
	}

	min := vals[0]
	for _, val := range vals[1:] {
		if val < min {
			min = val
		}
	}

	return min, nil
}

func max1(input int, vals ...int) (int, error) {

	max := input
	for _, val := range vals {
		if val > max {
			max = val
		}
	}

	return max, nil
}

func min1(input int, vals ...int) (int, error) {
	min := input
	for _, val := range vals {
		if val < min {
			min = val
		}
	}

	return min, nil
}