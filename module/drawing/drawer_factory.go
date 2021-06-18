package drawing

import (
	gingerinterface "github.com/HungTP-Play/ginger-go/ginger_interface"
)

func GetDrawer(spriteType SpriteType) gingerinterface.IDrawer {
	switch spriteType {
	case Square:
		return &RectDrawer{
			Name: "RectDrawer",
		}
	case Circle:
		return &CircelDrawer{
			Name: "CircleDrawer",
		}
	case Triangle:
		return &TriangleDrawer{
			Name: "TriangleDrawer",
		}
	case Rhombus:
		return &RhombusDrawer{
			Name: "RhombusDrawer",
		}
	default:
		return &RectDrawer{
			Name: "RectDrawer",
		}
	}
}
