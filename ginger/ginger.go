package ginger

import (
	"github.com/HungTP-Play/ginger-go/builder"
	"github.com/HungTP-Play/ginger-go/constant"
	"github.com/HungTP-Play/ginger-go/drawer/drawer"
	"github.com/HungTP-Play/ginger-go/util"
)

/// Draw Identicon image
func DrawIdenticon(idenInfo string, outputPath string, imgSize int, outputType constant.OutputType) error {
	identicon := builder.BuildIdenticon(idenInfo, outputType)
	util.PrepareOutput(outputPath)
	return drawer.DrawIdenticon(identicon, outputPath, imgSize)
}
