package main

import (
	"fmt"
	"sync"
)

var loadIconsOnce sync.Once

var icons map[string]string

func loadIcons() {
	icons = map[string]string{
		"spades.png":   "spades.png",
		"hearts.png":   "hearts.png",
		"diamonds.png": "diamonds.png",
		"clubs.png":    "clubs.png",
	}
}

func Icon(name string) string {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func main() {
	fmt.Println(Icon("hello"))
	fmt.Println(Icon("spades.png"))
}
