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

type Identicon struct {
	Name  string
	Hash  [16]byte
	Color [3]byte
	Size  IdenticonSize
	Grid  []byte
}
