package drawer

import (
	"image"
	"image/draw"
	"log"

	"github.com/HungTP-Play/ginger-go/constant"
	partdrawer "github.com/HungTP-Play/ginger-go/drawer/part_drawer"
	"github.com/HungTP-Play/ginger-go/model"
)

func DrawGithubIdenticon(identicon model.Identicon, outputDir string, imgSize int, padding int) draw.Image {
	drawSize := imgSize - padding*2
	img := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))

	// Draw backgournd
	DrawBackground(imgSize, img)

	// draw Identicon
	drawIdenticon(identicon, float64(drawSize), img, float64(padding))

	return img
}

// Draw all even value
func drawIdenticon(identicon model.Identicon, drawSize float64, img draw.Image, padding float64) {
	spriteType := constant.Type1
	drawer := partdrawer.GetPartDrawer(spriteType)
	sprites := buildSprites(identicon, drawSize, int(padding))
	log.Printf("Len sprites %v", len(sprites))
	for _, sprite := range sprites {
		drawer.DrawSprite(identicon.Color, img, sprite, 0)
	}
}

// Compute identicon sprite from it hash value, only odd value wiil be draw
//
// Step 1: Get 3 byte chunnk -> Reflex to 4 sysmmetry list
//
// ex: [1,2,3] -> [1,2,3,2,1]
//
// Step 2: Only the even value will be computed for graphic point (sprite)
func buildSprites(identicon model.Identicon, drawSize float64, padding int) []model.Sprite {
	w := drawSize / 5
	h := drawSize / 5

	hashList := []byte{}
	arrLen := len(identicon.Hash)

	for i := 0; i < arrLen && i+3 < arrLen-1; i += 3 {
		row := make([]byte, 5)
		copy(row, identicon.Hash[i:i+3])
		row[3] = row[1]
		row[4] = row[0]
		hashList = append(hashList, row...)
	}

	log.Printf("hashList len %v", len(hashList))

	sprites := []model.Sprite{}

	for index, value := range hashList {
		if value%2 == 0 {
			x := (index%5)*int(w) + padding
			y := (index/5)*int(h) + padding
			sprite := model.Sprite{
				TopLeft: model.Point{
					X: float64(x),
					Y: float64(y),
				},
				Width:  w,
				Height: h,
			}

			sprites = append(sprites, sprite)
		}
	}

	return sprites
}
