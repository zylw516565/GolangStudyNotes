package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"sync"
)

func main() {
	var filenames = []string{"fileA", "fileB", "fileC",}
	// makeThumbnails(filenames)
	// makeThumbnails2(filenames)
	// makeThumbnails3(filenames)
	// makeThumbnails4(filenames)
	makeThumbnails5(filenames)
}

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		ImageFile(f)
	}
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go ImageFile(f)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func (filename string) {
			time.Sleep(1 * time.Second)
			fmt.Printf("%s: sleeped %v seconds\n", filename, 1 * time.Second)
			ch <- struct{}{}
		}(f)

		// NOTE: incorrect!
		// go func () {
		// 	time.Sleep(1 * time.Second)
		// 	fmt.Printf("%s: sleeped %v seconds\n", f, 1 * time.Second)
		// 	ch <- struct{}{}
		// }()
	}

	for range filenames {
		<- ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
		_, err := ImageFile(f)
		errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: incorrect: goroutine leak!
		}
	}

	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err error
	}

	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

// makeThumbnails6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)

		// worker
		go func(f string) {
				defer wg.Done()
				thumb, err := ImageFile(f)
				if err != nil {
					log.Println(err)
					return
			}

			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

func ImageFile(filename string) (string, error) {
	time.Sleep(1 * time.Second)
	fmt.Printf("%s: sleeped %v seconds\n", filename, 1 * time.Second)
	return "", nil
}