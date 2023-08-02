package main

import (
	"os"
	"io"
	"fmt"
	"sort"
	"bufio"
	"unicode"
	"unicode/utf8"
)

func main() {
	var bTest bool
	fmt.Println(bTest)

	ages := make(map[string]int)
	fmt.Println(ages)

	ages2 := map[string]int{
		"alice": 31,
		"charlie": 34,
	}
	fmt.Println(ages2)

	ages3 := make(map[string]int)
	ages3["alice"] = 54
	ages3["charlie"] = 33
	fmt.Println(ages3)

	ages3["alice"] = 22
	fmt.Println(ages3["alice"])

	delete(ages3, "alice")
	fmt.Println(ages3)
	ages3["bob"] = ages3["bob"] + 1
	fmt.Println(ages3)
	ages3["bob"] += 1
	fmt.Println(ages3)
	ages3["bob"]++
	fmt.Println(ages3)
	// fmt.Println(&ages3["bob"])  //cannot take address of ages3["bob"]

	for name, age := range ages3 {
		fmt.Printf("%s\t%d\n", name, age)
	}

	mapNum := map[int]int{}
	for i := 	1; i <= 10; i++ {
		mapNum[i]=i
	}

	//每次遍历顺序随机
	for k, v := range mapNum {
		fmt.Printf("%d\t%d\n", k, v)
	}

	var nums []int
	for num := range mapNum {
		nums = append(nums, num)
	}

	sort.Ints(nums)
	for _, num := range nums {
		fmt.Printf("%d\t%d\n", num, mapNum[num])
	}

	ages4 := map[string]int{
		"alice": 31,
		"charlie": 34,
		"peter": 31,
		"jason": 34,
		"Lily": 31,
		"Lucy": 34,
	}

	// var names []string
	names := make([]string, 0, len(ages4))

	for name := range ages4 {
	names = append(names, name)
	} 
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages4[name])
	}

	var ages5 map[string]int
	fmt.Println(ages5 == nil)
	fmt.Println(len(ages5) == 0)
	// ages5["a"] = 1    //panic: assignment to entry in nil map
	ages6 := make(map[string]int, 10)
	fmt.Println(ages6 == nil)
	fmt.Println(len(ages6) == 0)
	ages6["a"] = 1

	ages7 := map[string]int{
		"alice": 31,
		"charlie": 34,
		"peter": 31,
		"jason": 34,
		"Lily": 31,
		"Lucy": 34,
	}

	age, ok := ages7["bob"]
	if !ok {
		fmt.Println("bob is not in map", age)
	}

	if _, ok := ages["bob"]; !ok { /* ... */ }

	fmt.Println(equal(ages7, ages7))
	fmt.Println(equal(ages6, ages7))

	// dedup()
	var list []byte
	for i := byte(0); i < 128; i++ {
		list = append(list, i)
	}
	fmt.Println(k2(list))

	fmt.Println("\x00\x01")
	fmt.Println("\a")
	fmt.Printf("%s", "aaaaa\v")
	fmt.Printf("%s", "aaaaa\f")
	fmt.Printf("%s", "aaaaa\r")

	testASCII();

	var str = []string{"abcccccdef", "abcdddddd", "aaaabbbb"}
	fmt.Println(str)
	fmt.Println(k(str))

	var str2 = []string{"a", "b", "c"}
	var str3 = []string{"a", "b", "c"}
	var str4 = []string{"a", "c", "c"}
	Add(str)
	Add(str2)
	Add(str3)
	Add(str4)
	fmt.Println(Count(str), Count(str2), Count(str3), Count(str4))

	// charCount()

	graph := make(map[string]map[string]bool)
	fmt.Printf("graph type is:%T\n", graph)

	charCount2()
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, v := range x {
		if _, ok := y[k]; !ok {
			return false
		}else {
			if v != y[k] {
				return  false
			}
		}
	}

	return true
}

func dedup() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
		seen[line] = true
		fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func testASCII() {
	f, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	output := bufio.NewWriter(f)
	output.WriteString("aaaaa\v")
	output.WriteString("aaaaa\f")
	output.WriteString("aaaaa")
	err = output.Flush()
	if err != nil {
		fmt.Fprintf(os.Stderr, "output.Flush: %v\n", err)
		os.Exit(1)
	}
}

var m = make(map[string]int)

func k(list []string) string {return fmt.Sprintf("%q", list)}
func Add(list []string) {m[k(list)]++}
func Count(list []string) int {return m[k(list)]}

func k2(list []byte) string {return fmt.Sprintf("%q", list)}

func charCount() {
	counts := make(map[rune]int) // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0 // count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")

	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")

	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

const (
	LETTER = iota + 1
	DIGITAL
	COUNT_MAX
)

func charCount2() {
	counts := make(map[rune]int) // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0 // count of invalid UTF-8 characters
	var letters, digits int
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		if unicode.IsLetter(r) {
			letters++
		}

		if unicode.IsNumber(r) {
			digits++
		}
	}
	fmt.Printf("letters\tcount:%d\n", letters)
	fmt.Printf("digits\tcount:%d\n", digits)

	fmt.Printf("rune\tcount\n")

	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")

	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}