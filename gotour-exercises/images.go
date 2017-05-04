package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Height, Width int
}

func (I Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, I.Height, I.Width)
}

func (I Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (I Image) At(x, y int) color.Color {
	c := uint8(x ^ y)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
