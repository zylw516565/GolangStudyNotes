package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
	"crypto/sha256"
	"crypto/sha512"
)

var algo = flag.String("a", "sha256", "specific hash algorithm")

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		switch  strings.ToLower(*algo) {
		case "sha384":
			fmt.Printf("%v's hash code:\n%x\n", arg, sha512.Sum384([]byte(arg)))
		case "sha512":
			fmt.Printf("%v's hash code:\n%x\n", arg, sha512.Sum512([]byte(arg)))
		case "sha256":
			fmt.Printf("%v's hash code:\n%x\n", arg, sha256.Sum256([]byte(arg)))
		default:
			fmt.Println("invalid hash algorithm")
			os.Exit(1)
		}
	}
}