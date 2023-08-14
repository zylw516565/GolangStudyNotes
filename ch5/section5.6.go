package main

import (
	"os"
	"fmt"
	"sort"
	"strings"
)

var prereqs = map[string][]string {
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
	"data structures",
	"formal languages",
	"computer organization",
	},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func squares() func () int {
	var x int
	return func () int {
		x++
		return x * 2
	}
}

func main () {
	fmt.Println(strings.Map(func (r rune) rune {return r + 3}, "HAL-9000")) // "IBM.:111"
	fmt.Println()

	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println()
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

	testForRange()
	// testForRange2()
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

// func topoSort2(m map[string][]string) []string {
// 	var order []string
// 	seen := make(map[string]bool)

// 	visitAll := func(items []string) {
// 		for _, item := range items {
// 			if !seen[item] {
// 				seen[item] = true
// 				visitAll(m[item])    //undefined: visitAll
// 				order = append(order, item)
// 			}
// 		}
// 	}

// 	var keys []string
// 	for key := range m {
// 		keys = append(keys, key)
// 	}

// 	sort.Strings(keys)
// 	visitAll(keys)
// 	return order
// }


func testForRange() {
	var rmdirs [] func()

	var tempDirs = []string{"/home/1", "/home/2", "/home/3",}

	for _, dir := range tempDirs {
		dir := dir		// NOTE: necessary!
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func () {
				os.RemoveAll(dir)
			})
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}
}

func testForRange2() {
	var rmdirs [] func()

	var tempDirs = []string{"/home/1", "/home/2", "/home/3",}
	dirs := tempDirs

	for i := 0; i < len(dirs); i++ {
		os.MkdirAll(dirs[i], 0755)
		rmdirs = append(rmdirs, func () {
				os.RemoveAll(dirs[i])    // NOTE: incorrect! Will panic when rmdir. i = 3;
			})
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}
}