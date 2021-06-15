package builder

import (
	"crypto/sha512"

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

var NumSprite = map[model.IdenticonSize]int{
	model.Size3x3: 9,
	model.Size4x4: 16,
	model.Size5x5: 25,
	model.Size6x6: 36,
	model.Size7x7: 49,
	model.Size8x8: 64,
	model.Size9x9: 81,
}

var ChunkSize = map[model.IdenticonSize]int{
	model.Size3x3: 3,
	model.Size4x4: 4,
	model.Size5x5: 5,
	model.Size6x6: 6,
	model.Size7x7: 7,
	model.Size8x8: 8,
	model.Size9x9: 9,
}

// Build Identicon struct from input string
func IdenticonInit(input []byte, iconSize model.IdenticonSize) model.Identicon {
	checksum := sha512.Sum512(input)
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
	grid := []byte{}
	step := StepMap[indenticon.Size]
	numStep := NumSprite[indenticon.Size]
	chunkSize := ChunkSize[indenticon.Size]
	for i := 0; i < int(numStep/2); i++ {
		chunk := make([]byte, chunkSize)
		copy(chunk, indenticon.Hash[i:i+step])
		chunk = ReflexArr(chunk, chunkSize)
		grid = append(grid, chunk...)
	}
	indenticon.Grid = grid
	return indenticon
}

// Turn arr to refexed arr
//
// ex: [1,2,3] -> [1,2,1], [1,2,3,4] -> [1,2,2,1]
func ReflexArr(arr []byte, n int) []byte {
	for i := 0; i < int(n/2); i++ {
		arr[n-i-1] = arr[i]
	}
	return arr
}

func BuildIdenticon(input []byte, iconSize model.IdenticonSize) model.Identicon {
	identicon := IdenticonInit(input, iconSize)
	identicon = BuildGrid(identicon)
	return identicon
}
