package ginger

import (
	"errors"
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path"

	"github.com/HungTP-Play/ginger-go/builder"
	"github.com/HungTP-Play/ginger-go/constant"
	"github.com/HungTP-Play/ginger-go/drawer/drawer"
	"github.com/HungTP-Play/ginger-go/model"
	"github.com/HungTP-Play/ginger-go/util"
	"github.com/llgcode/draw2d/draw2dimg"
)

func getExt(outputType constant.OutputType) string {
	ext := ""
	if outputType == constant.SVG {
		ext = ".svg"
	} else if outputType == constant.PNG {
		ext = ".png"
	} else {
		ext = ".jpg"
	}
	return ext
}

/// Draw Identicon image
func DrawIdenticon(idenInfo string, outputPath string, imgSize int, padding int, outputType constant.OutputType) (*model.Identicon, error) {
	identicon := builder.BuildIdenticon(idenInfo, outputType)
	util.PrepareOutput(outputPath)
	img := drawer.DrawIdenticon(identicon, outputPath, imgSize, padding)

	ext := getExt(outputType)

	outputFilePath := path.Join(outputPath, fmt.Sprintf("%v%v", identicon.IdenInfo, ext))
	log.Printf("Output path: %v", outputFilePath)
	var err error = nil
	if outputType == constant.SVG {
		return nil, errors.New("SVG is not supported at this time")
	} else if outputType == constant.PNG {
		err = draw2dimg.SaveToPngFile(outputFilePath, img)
	} else {
		f, err := os.Create(outputFilePath)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		jpeg.Encode(f, img, nil)
	}
	return &identicon, err
}

// Draw Github Identicon
func DrawGithubIdenticon(idenInfo string, outputPath string, imgSize int, padding int, outputType constant.OutputType) (*model.Identicon, error) {
	identicon := builder.BuildIdenticon(idenInfo, outputType)
	util.PrepareOutput(outputPath)
	img := drawer.DrawGithubIdenticon(identicon, outputPath, imgSize, padding)

	ext := getExt(outputType)

	outputFilePath := path.Join(outputPath, fmt.Sprintf("%v%v", identicon.IdenInfo, ext))
	log.Printf("Output path: %v", outputFilePath)
	var err error = nil
	if outputType == constant.SVG {
		return nil, errors.New("SVG is not supported at this time")
	} else if outputType == constant.PNG {
		err = draw2dimg.SaveToPngFile(outputFilePath, img)
	} else {
		f, err := os.Create(outputFilePath)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		jpeg.Encode(f, img, nil)
	}
	return &identicon, err
}
