package main

import (
	"os"
	"fmt"
	"sort"
	"time"
	"text/tabwriter"
)

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

// i, j are indices of sequence elements
func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i] , s[j] = s[j], s[i]
}

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

var tracks = []*Track{
	{"Go",         "Delilah",        "From the Roots Up", 2012, length("3m38s")},
	{"Go",         "Moby",           "Moby",              1992, length("3m37s")},
	{"Go Ahead",   "Alicia Keys",    "As I Am",           2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash",             2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}

	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}

	tw.Flush() // calculate column widths and print table
}

//****************************
type byArtists []*Track

func (s byArtists) Len() int {
	return len(s)
}

// i, j are indices of sequence elements
func (s byArtists) Less(i, j int) bool {
	return s[i].Artist < s[j].Artist
}

func (s byArtists) Swap(i, j int) {
	s[i] , s[j] = s[j], s[i]
}

//****************************
type byYear []*Track

func (s byYear) Len() int {
	return len(s)
}

// i, j are indices of sequence elements
func (s byYear) Less(i, j int) bool {
	return s[i].Year < s[j].Year
}

func (s byYear) Swap(i, j int) {
	s[i] , s[j] = s[j], s[i]
}

//****************************
type customSort struct {
	t []*Track
	less func (x, y *Track) bool
}

func (c customSort) Len() int {return len(c.t)}
func (c customSort) Less(i, j int) bool { return c.less(c.t[i], c.t[j])}
func (c customSort) Swap(i, j int) {c.t[i], c.t[j] = c.t[j], c.t[i]}

func main() {
	s := StringSlice{"hello", "world", "ni", "hao"}
	fmt.Println(s)
	sort.Sort(s)
	fmt.Println(s)

	fmt.Println("****************************")
	fmt.Println("Before sort !!!")
	fmt.Println("****************************")
	printTracks(tracks)

	fmt.Println("****************************")
	fmt.Println("After sort !!!")
	fmt.Println("****************************")
	sort.Sort(byArtists(tracks))
	printTracks(tracks)

	fmt.Println("****************************")
	fmt.Println("Reverse sort !!!")
	fmt.Println("****************************")
	sort.Sort(sort.Reverse(byArtists(tracks)))
	printTracks(tracks)

	fmt.Println("****************************")
	fmt.Println("Sort by year !!!")
	fmt.Println("****************************")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("****************************")
	fmt.Println("customSort !!!")
	fmt.Println("****************************")
	sort.Sort(customSort{tracks, func (x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}

		if x.Artist != y.Artist {
			return x.Artist < y.Artist
		}

		if x.Year != y.Year {
			return x.Year < y.Year
		}

		return false
	}})
	printTracks(tracks)

	fmt.Println()
	fmt.Println("****************************")
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println(values) // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values) // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"
}