package model

import (
	"image/color"

	"github.com/HungTP-Play/ginger-go/constant"
)

type Identicon struct {
	IdenInfo          string
	Hash              [16]byte
	Color             color.Color
	CenterSpriteIndex int
	SideSpriteIndex   int
	CornerSpriteIndex int
	OutputType        constant.OutputType
	StartSideIndex    int
	StartCornerIndex  int
}
