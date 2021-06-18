package drawing

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"path/filepath"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

func DrawIdenticon(identicon model.Identicon, outputDir string, spriteType SpriteType) error {
	// TODO: Need to do re-arange here
	// We create our default image containing a 250x250 rectangle
	var img = image.NewRGBA(image.Rect(0, 0, identicon.ImgSize, identicon.ImgSize))

	fileOutputPath := filepath.Join(outputDir, fmt.Sprintf("%v.png", identicon.Name))

	os.MkdirAll(outputDir, 0700)

	whiteCol := color.RGBA{255, 255, 255, 255}

	// Draw back ground
	DrawRect(
		img,
		whiteCol,
		0,
		0,
		float64(identicon.ImgSize),
		float64(identicon.ImgSize),
	)

	// We retrieve the color from the color property on the identicon
	col := color.RGBA{identicon.Color[0], identicon.Color[1], identicon.Color[2], 255}

	// Loop over the pixelmap and call the rect function with the img, color and the dimensions
	drawer := GetDrawer(spriteType)
	for _, pixel := range identicon.DrawingPoints {
		drawer.Draw(
			img,
			col,
			float64(pixel.TopLeft.X),
			float64(pixel.TopLeft.Y),
			float64(pixel.BottomRight.X),
			float64(pixel.BottomRight.Y),
		)
	}
	log.Printf("Output path:%v", fileOutputPath)
	return draw2dimg.SaveToPngFile(fileOutputPath, img)
}
