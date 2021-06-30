package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

type Pixel struct {
	R int
	G int
	B int
}

func distance(pixel1, pixel2 Pixel) int {
	return (pixel1.R-pixel2.R)*(pixel1.R-pixel2.R) +
		(pixel1.G-pixel2.G)*(pixel1.G-pixel2.G) +
		(pixel1.B-pixel2.B)*(pixel1.B-pixel2.B)
}

func main() {
	file, _ := os.Open("image.png")
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	pixels := make([]Pixel, img.Bounds().Max.Y*img.Bounds().Max.X)

	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			pixels[img.Bounds().Max.X*y+x] = Pixel{R: int(r / 257), G: int(g / 257), B: int(b / 257)}
			// fmt.Printf("%v", color)
		}
	}
	fmt.Println(pixels)
}
