package builder

import (
	"crypto/md5"
	"image/color"

	"github.com/HungTP-Play/ginger-go/constant"
	"github.com/HungTP-Play/ginger-go/model"
)

func BuildIdenticon(idenInfo string, outputType constant.OutputType) model.Identicon {
	hash := md5.Sum([]byte(idenInfo))
	color := PickColor(hash)
	centerIndex, sideIndex, cornerIndex := PickSpriteIndex(hash)
	startSideIndex, startCornerIndex := PickStartIndices(hash)
	return model.Identicon{
		IdenInfo:          idenInfo,
		Hash:              hash,
		OutputType:        outputType,
		CenterSpriteIndex: centerIndex,
		SideSpriteIndex:   sideIndex,
		CornerSpriteIndex: cornerIndex,
		Color:             color,
		StartSideIndex:    startSideIndex,
		StartCornerIndex:  startCornerIndex,
	}
}

// Pick first 3 value of has
func PickColor(hash [16]byte) color.Color {
	red := hash[0]
	green := hash[1]
	blue := hash[2]
	color := color.RGBA{red, green, blue, 255}
	return color
}

func PickSpriteIndex(hash [16]byte) (centerIndex int, sideIndex int, cornerIndex int) {
	centerIndex = int((hash[3] & 192) >> 6)
	sideIndex = int((hash[3] & 60) >> 2)
	cornerIndex = int(((hash[3] & 3 << 6) | (hash[4] & 192 >> 2)) >> 4)
	return centerIndex, sideIndex, cornerIndex
}

func PickStartIndices(hash [16]byte) (startSideIndex int, StartCornerIndex int) {
	startSideIndex = int((hash[4] & 60) >> 4)
	StartCornerIndex = int((hash[4] & 12 >> 2))
	return startSideIndex, StartCornerIndex
}
