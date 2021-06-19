package partdrawer

import (
	"image/color"
	"image/draw"

	"github.com/HungTP-Play/ginger-go/model"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Type9Drawer struct {
	Name string
}

func (p *Type9Drawer) DrawSprite(color color.Color, image draw.Image, sprite model.Sprite) {
	x1, y1 := sprite.TopLeft.X, sprite.TopLeft.Y
	x1p4, x3p4 := x1+sprite.Width*0.25, x1+sprite.Width*0.75
	y1p4, y3p4 := y1+sprite.Height*0.25, y1+sprite.Height*0.75
	gc := draw2dimg.NewGraphicContext(image)

	// TopLeft -> TopRight
	gc.MoveTo(x1p4, y1p4)
	gc.LineTo(x3p4, y1p4)

	// TopRight -> BottomRight
	gc.MoveTo(x3p4, y1p4)
	gc.LineTo(x3p4, y3p4)

	// BottomRight -> BottomLeft
	gc.MoveTo(x3p4, y3p4)
	gc.LineTo(x1p4, y3p4)

	// BottomLeft -> TopLeft
	gc.MoveTo(x1p4, y3p4)
	gc.LineTo(x1p4, y1p4)

	gc.SetLineWidth(0)

	// Fill
	gc.FillStroke()
}
