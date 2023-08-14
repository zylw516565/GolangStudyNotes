package somepackage

type T struct {
	a, b int
}

type points struct {
	X, Y int
}

type circles struct {
	points
	Radius int
}

type White struct {
	circles
	Spokes int
}