package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	file, _ := os.Open("image.png")
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	pixels := make([]color.Color, img.Bounds().Max.Y*img.Bounds().Max.X)

	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			color := img.At(x, y)
			pixels[img.Bounds().Max.X*y+x] = color
			// fmt.Printf("%v", color)
		}
	}
	fmt.Println(pixels[:img.Bounds().Max.X])
}
