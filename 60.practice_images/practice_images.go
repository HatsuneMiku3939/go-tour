package main

import (
//    "code.google.com/p/go-tour/pic"
    "image/color"
    "image"
)

type Image struct {
    width int
    height int
    pixels []color.RGBA
}

func NewImage(width int, height int) *Image {
    img := Image { }

    img.width = width
    img.height = height
    img.pixels = make([]color.RGBA, width * height)

    for i := range img.pixels {
        img.pixels[i] = color.RGBA {uint8(i % width), uint8(i / width), 255, 255}
    }

    return &img
}

func (img *Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
    return image.Rectangle {image.Point {0, 0}, image.Point {img.width, img.height}}
}

func (img *Image) At(x, y int) color.Color {
    return img.pixels[x * y]
}

func main() {
    m := NewImage(255, 255)
    ShowImage(m)
}
