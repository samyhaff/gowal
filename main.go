package main

import (
	"fmt"
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
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			color := img.At(x, y)
			fmt.Printf("%v", color)
		}
	}
}
