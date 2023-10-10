package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
	"sync"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

var done = make(chan struct{})
func cancelled() bool {
	select {
		case <-done:
			return true
		default:
			return false
	}
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	// Print the results.
	var nfiles, nbytes int64
loop:
	for {
		select {
			case size, ok := <-fileSizes:
				if !ok {
					break loop
				}
				nfiles++
				nbytes += size
			case <-tick:
				printDiskUsage(nfiles, nbytes)
			case <-done:
				// Drain fileSizes to allow existing goroutines to finish.
				for range fileSizes {
					// Do nothing.
				}
				panic(1)
				// return
			}
	}

	printDiskUsage(nfiles, nbytes)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	var sema = make(chan struct{}, 20)

	select {
		case sema <- struct{}{}:
		case <-done:
			return nil
	}
	defer func() {<- sema}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}