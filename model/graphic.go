package model

type Point struct {
	X, Y int
}

type DrawingPoint struct {
	TopLeft     Point
	BottomRight Point
}
