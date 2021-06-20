package drawer

import (
	"image"
	"image/color"
	"image/draw"
	"sync"

	"github.com/HungTP-Play/ginger-go/constant"
	partdrawer "github.com/HungTP-Play/ginger-go/drawer/part_drawer"
	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

// Draw full identicon with desired size and then save into outputDir
//
// Image will use idenInfo as it name
//
// ex: idenInfo is "stoicmeme" and use PNG format, then image will be drawed to outputDir/storicmeme.png
func DrawIdenticon(identicon model.Identicon, outputDir string, imgSize int, padding int) draw.Image {
	drawSize := imgSize - padding*2
	img := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))
	var wg sync.WaitGroup

	// Draw backgournd
	DrawBackground(imgSize, img)

	// Draw center part
	wg.Add(1)
	go func() {
		defer wg.Done()
		drawCenter(identicon, drawSize, img, float64(padding))
	}()

	// Draw side parts
	wg.Add(1)
	go func() {
		defer wg.Done()
		drawSides(identicon, drawSize, img, float64(padding))
	}()

	// Draw corner parts
	wg.Add(1)
	go func() {
		defer wg.Done()
		drawCorners(identicon, drawSize, img, float64(padding))
	}()

	wg.Wait()
	return img
}

// Draw background white
func DrawBackground(imgSize int, image draw.Image) {
	gc := draw2dimg.NewGraphicContext(image)
	gc.SetFillColor(color.White)
	gc.MoveTo(0, 0)
	gc.LineTo(float64(imgSize), 0)
	gc.MoveTo(float64(imgSize), 0)
	gc.LineTo(float64(imgSize), float64(imgSize))
	gc.MoveTo(float64(imgSize), float64(imgSize))
	gc.LineTo(0, float64(imgSize))
	gc.MoveTo(0, float64(imgSize))
	gc.LineTo(0, 0)
	gc.SetLineWidth(0)
	gc.FillStroke()
}

// Draw center part
func drawCenter(identicon model.Identicon, drawSize int, image draw.Image, padding float64) {
	centerSprite := model.Sprite{
		TopLeft: model.Point{
			X: float64(drawSize)/3 + padding,
			Y: float64(drawSize)/3 + padding,
		},
		Width:  float64(drawSize) / 3,
		Height: float64(drawSize) / 3,
	}

	drawer := partdrawer.GetPartDrawer(constant.CenterTypes[identicon.CenterSpriteIndex])

	drawer.DrawSprite(identicon.Color, image, centerSprite, 0)
}

// Draw side parts
func drawSides(identicon model.Identicon, drawSize int, image draw.Image, padding float64) {
	w := drawSize / 3
	h := drawSize / 3

	// Top
	top := model.Sprite{
		TopLeft: model.Point{
			X: float64(drawSize)/3 + padding,
			Y: 0 + padding,
		},
		Width:  float64(w),
		Height: float64(h),
	}

	// Right
	right := model.Sprite{
		TopLeft: model.Point{
			X: float64(drawSize)/3*2 + padding,
			Y: float64(drawSize)/3 + padding,
		},
		Width:  float64(w),
		Height: float64(h),
	}

	// Bottom
	bottom := model.Sprite{
		TopLeft: model.Point{
			X: float64(drawSize)/3 + padding,
			Y: float64(drawSize)/3*2 + padding,
		},
		Width:  float64(w),
		Height: float64(h),
	}

	// Left
	left := model.Sprite{
		TopLeft: model.Point{
			X: 0 + padding,
			Y: float64(drawSize)/3 + padding,
		},
		Width:  float64(w),
		Height: float64(h),
	}

	sprites := [4]model.Sprite{top, right, bottom, left}

	drawer := partdrawer.GetPartDrawer(constant.AllTypes[identicon.SideSpriteIndex])

	for index, sprite := range sprites {
		rotation := CalRotation(index, identicon.StartSideIndex)
		drawer.DrawSprite(identicon.Color, image, sprite, rotation)
	}
}

// Draw corner parts
func drawCorners(identicon model.Identicon, drawSize int, image draw.Image, padding float64) {
	w := drawSize / 3
	h := drawSize / 3

	// TopLeft
	topLeft := model.Sprite{
		TopLeft: model.Point{
			X: 0 + padding,
			Y: 0 + padding,
		},
		Width:  float64(w),
		Height: float64(h),
	}

	// TopRight
	topRight := model.Sprite{
		TopLeft: model.Point{
			X: float64(drawSize)/3*2 + padding,
			Y: 0 + padding,
		},
		Width:  float64(w),
		Height: float64(h),
	}

	// BottomRight
	bottomRight := model.Sprite{
		TopLeft: model.Point{
			X: float64(drawSize)/3*2 + padding,
			Y: float64(drawSize)/3*2 + padding,
		},
		Width:  float64(w),
		Height: float64(h),
	}

	// BottomLeft
	bottomLeft := model.Sprite{
		TopLeft: model.Point{
			X: 0 + padding,
			Y: float64(drawSize)/3*2 + padding,
		},
		Width:  float64(w),
		Height: float64(h),
	}

	sprites := [4]model.Sprite{topLeft, topRight, bottomRight, bottomLeft}

	drawer := partdrawer.GetPartDrawer(constant.AllTypes[identicon.CornerSpriteIndex])

	for index, sprite := range sprites {
		rotation := CalRotation(index, identicon.StartCornerIndex)
		drawer.DrawSprite(identicon.Color, image, sprite, rotation)
	}
}

func CalRotation(index int, startIndex int) int {
	if index == startIndex {
		return 0
	}
	diff := index - startIndex
	if diff < 0 {
		diff = -diff
		return 4 - diff
	} else {
		return diff
	}
}
