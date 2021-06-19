package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Type2Drawer struct {
	Name string
}

func (p *Type2Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int) {
	x2, y2 := sprite.TopLeft.X+sprite.Width, sprite.TopLeft.Y+sprite.Height
	x1, y1 := sprite.TopLeft.X, sprite.TopLeft.Y
	gc := draw2dimg.NewGraphicContext(image)

	if rotation == 0 {
		// TopLeft -> TopRight
		gc.MoveTo(x1, y1)
		gc.LineTo(x2, y1)

		// TopRight -> BottomLeft
		gc.MoveTo(x2, y1)
		gc.LineTo(x1, y2)

		// BottomLeft -> TopLeft
		gc.MoveTo(x1, y2)
		gc.LineTo(x1, y1)
	} else if rotation == 1 {
		// TopLeft -> TopRight
		gc.MoveTo(x1, y1)
		gc.LineTo(x2, y1)

		// TopRight -> BottomRight
		gc.MoveTo(x2, y1)
		gc.LineTo(x2, y2)

		// BottomRight -> TopLeft
		gc.MoveTo(x2, y2)
		gc.LineTo(x1, y1)
	} else if rotation == 2 {
		// TopRight -> BottomRight
		gc.MoveTo(x2, y1)
		gc.LineTo(x2, y2)

		// BottomRight -> BottomLeft
		gc.MoveTo(x2, y2)
		gc.LineTo(x1, y2)

		// BottomLeft -> TopRight
		gc.MoveTo(x1, y2)
		gc.LineTo(x2, y1)
	} else {
		// TopLeft -> BottomRight
		gc.MoveTo(x1, y1)
		gc.LineTo(x2, y2)

		// BottomRight -> BottomLeft
		gc.MoveTo(x2, y2)
		gc.LineTo(x1, y2)

		// BottomLeft -> TopLeft
		gc.MoveTo(x1, y2)
		gc.LineTo(x1, y1)
	}

	gc.SetLineWidth(0)

	// Fill
	gc.FillStroke()
}
