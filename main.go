package main

import (
	"errors"
	"fmt"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

const COLORS_NUMBER = 16

type Pixel struct {
	R int
	G int
	B int
}

type Cluster struct {
	centroid Pixel
	members  []Pixel
}

func distance(pixel1, pixel2 Pixel) int {
	return (pixel1.R-pixel2.R)*(pixel1.R-pixel2.R) +
		(pixel1.G-pixel2.G)*(pixel1.G-pixel2.G) +
		(pixel1.B-pixel2.B)*(pixel1.B-pixel2.B)
}

func closestCentroid(pixel Pixel, clusters [COLORS_NUMBER]Cluster) int {
	var min_i int
	min_dist := distance(pixel, clusters[0].centroid)

	for i := 0; i < COLORS_NUMBER; i++ {
		dist := distance(pixel, clusters[i].centroid)
		if dist < min_dist {
			min_dist = dist
			min_i = i
		}
	}
	return min_i
}

func mean(pixels []Pixel) Pixel {
	r, g, b := 0, 0, 0
	for _, pixel := range pixels {
		r += pixel.R
		g += pixel.G
		b += pixel.B
	}
	r, g, b = r/len(pixels), g/len(pixels), b/len(pixels)
	return Pixel{R: r, G: g, B: b}
}

func assign(pixels []Pixel, clusters [COLORS_NUMBER]Cluster) [COLORS_NUMBER]Cluster {
	for _, pixel := range pixels {
		clusters[closestCentroid(pixel, clusters)].members = append(clusters[closestCentroid(pixel, clusters)].members, pixel)
	}

	return clusters
}

func initialize(pixels []Pixel) [COLORS_NUMBER]Cluster {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var clusters [COLORS_NUMBER]Cluster
	// start by choosing random centroids
	for i := 0; i < COLORS_NUMBER; i++ {
		clusters[i] = Cluster{centroid: pixels[r.Intn(len(pixels))], members: []Pixel{}}
	}

	// get the members of each cluster
	clusters = assign(pixels, clusters)

	return clusters
}

func iterate(clusters [COLORS_NUMBER]Cluster, pixels []Pixel) ([COLORS_NUMBER]Cluster, bool) {
	changed := false

	for i := 0; i < COLORS_NUMBER; i++ {
		var new_centroid Pixel
		// centroid becomes the mean of the cluster members
		if len(clusters[i].members) == 0 {
			continue
		}
		new_centroid = mean(clusters[i].members)
		if clusters[i].centroid != new_centroid {
			changed = true
		}
		clusters[i].centroid = new_centroid
	}

	if changed {
		for i := 0; i < COLORS_NUMBER; i++ {
			clusters[i].members = []Pixel{}
		}

		clusters = assign(pixels, clusters)
	}

	return clusters, changed
}

func main() {
	// read arguments
	if len(os.Args) < 2 {
		err := errors.New("No image was given")
		log.Fatal(err)
	}
	path := os.Args[1]

	// open image
	file, _ := os.Open(path)
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// read image into array
	pixels := make([]Pixel, img.Bounds().Max.Y*img.Bounds().Max.X)
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			pixels[img.Bounds().Max.X*y+x] = Pixel{R: int(r / 257), G: int(g / 257), B: int(b / 257)}
			// fmt.Printf("%v", color)
		}
	}

	// main algorithm

	clusters := initialize(pixels)

	changed := true
	for changed {
		clusters, changed = iterate(clusters, pixels)
	}

	for _, cluster := range clusters {
		fmt.Println(fmt.Sprintf("#%02X%02X%02X", cluster.centroid.R, cluster.centroid.G, cluster.centroid.B))
	}
}
