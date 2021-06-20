package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Type1Drawer struct {
	Name string
}

func (p *Type1Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite, rotation int) {
	x2, y2 := sprite.TopLeft.X+sprite.Width, sprite.TopLeft.Y+sprite.Height
	x1, y1 := sprite.TopLeft.X, sprite.TopLeft.Y
	gc := draw2dimg.NewGraphicContext(image)
	gc.SetFillColor(color)

	// TopLeft -> TopRight
	gc.MoveTo(x1, y1)
	gc.LineTo(x2, y1)

	// TopRight -> BottomRight
	gc.LineTo(x2, y2)

	// BottomRight -> BottomLeft
	gc.LineTo(x1, y2)

	// BottomLeft -> TopLeft
	gc.LineTo(x1, y1)

	gc.SetLineWidth(0)

	// Fill
	gc.FillStroke()
}
