package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	Width, Height int
	clr uint8	
}

func (img *Image) ColorModel() color.Model{
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle{
	return image.Rect(0,0, img.Width, img.Height)
}

func (img *Image) At(x, y int) color.Color{
	return color.RGBA{img.clr + uint8(x), img.clr + uint8(y), 255, 255}
}

func main() {
	m := Image{50, 80, 120}
	pic.ShowImage(&m)
}
