package builder

import (
	"crypto/md5"

	"github.com/HungTP-Play/ginger-go/model"
)

var StepMap = map[model.IdenticonSize]int{
	model.Size3x3: 2,
	model.Size4x4: 2,
	model.Size5x5: 3,
	model.Size6x6: 3,
	model.Size7x7: 4,
	model.Size8x8: 4,
	model.Size9x9: 5,
}

var NeedPivotMap = map[model.IdenticonSize]bool{
	model.Size3x3: true,
	model.Size4x4: false,
	model.Size5x5: true,
	model.Size6x6: false,
	model.Size7x7: true,
	model.Size8x8: false,
	model.Size9x9: true,
}

// Build Identicon struct from input string
func IdenticonBuilder(input []byte, iconSize model.IdenticonSize) model.Identicon {
	checksum := md5.Sum(input)
	color := PickCorlor(checksum[:])

	return model.Identicon{
		Name:  string(input),
		Hash:  checksum,
		Color: color,
		Size:  iconSize,
	}
}

// Get first 3 byte to define the color RGBokm
func PickCorlor(input []byte) [3]byte {
	color := [3]byte{}
	copy(color[:], input[:3])
	return color
}

func BuildGrid(indenticon model.Identicon) model.Identicon {
	// grid := []byte{}
	step := StepMap[indenticon.Size]
	// needPivot := NeedPivotMap[indenticon.Size]
	if step == 5 {
		// TODO: implement this special case
	}

	return indenticon
}
