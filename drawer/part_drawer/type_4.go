package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Type4Drawer struct {
	Name string
}

func (p *Type4Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int) {
	x2, y2 := sprite.TopLeft.X+sprite.Width, sprite.TopLeft.Y+sprite.Height
	x1, y1 := sprite.TopLeft.X, sprite.TopLeft.Y
	xCenter, yCenter := x1+sprite.Width/2, y1+sprite.Height/2
	gc := draw2dimg.NewGraphicContext(image)
	gc.SetFillColor(color)

	if rotation == 0 {
		// TopLeft -> TopCenter
		gc.MoveTo(x1, y1)
		gc.LineTo(xCenter, y1)

		// TopCenter -> BottomCenter
		gc.LineTo(xCenter, y2)

		// BottomCenter -> BottomLLeft
		gc.LineTo(x1, y2)

		// BottomLeft -> TopLeft
		gc.LineTo(x1, y1)
	} else if rotation == 1 {
		// TopLeft -> TopRight
		gc.MoveTo(x1, y1)
		gc.LineTo(x2, y1)

		// TopRight -> RightCenter
		gc.LineTo(x2, yCenter)

		// RightCenter -> LeftCenter
		gc.LineTo(x1, yCenter)

		// LeftCenter -> TopLeft
		gc.LineTo(x1, y1)
	} else if rotation == 2 {
		// TopCenter -> TopRight
		gc.MoveTo(xCenter, y1)
		gc.LineTo(x2, y1)

		// TopRight -> BottomRight
		gc.LineTo(x2, y2)

		// BottomRight -> BottomCenter
		gc.LineTo(xCenter, y2)

		// BottomCenter -> TopCenter
		gc.LineTo(xCenter, y1)
	} else {
		// RightCeter -> BottomRight
		gc.MoveTo(x2, yCenter)
		gc.LineTo(x2, y2)

		// BottomRight -> BottomLeft
		gc.LineTo(x1, y2)

		// BottomLeft -> LeftCenter
		gc.LineTo(x1, yCenter)

		// LeftCenter -> RightCenter
		gc.LineTo(x2, yCenter)
	}

	gc.SetLineWidth(0)

	// Fill
	gc.FillStroke()
}
