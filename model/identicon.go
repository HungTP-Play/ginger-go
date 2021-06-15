package model

type IdenticonSize int

const (
	Size3x3 IdenticonSize = iota
	Size4x4
	Size5x5
	Size6x6
	Size7x7
	Size8x8
	Size9x9
)

type GridPoint struct {
	Value byte
	Index int
}
type Identicon struct {
	Name          string
	Hash          [64]byte
	Color         [3]byte
	Size          IdenticonSize
	Grid          []byte
	GridPoints    []GridPoint
	DrawingPoints []DrawingPoint
	ImgSize       int
	NeedRearange  bool
}
