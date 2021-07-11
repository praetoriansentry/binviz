package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
)

const (
	IMAGE_WIDTH  = 2048
	IMAGE_HEIGHT = 2048
	MAX_DIM      = (255 << 8) | 255
)

type (
	Coord struct {
		X int
		Y int
	}
)

func main() {
	log.Println("Test")
	bin, err := ioutil.ReadFile("main")
	if err != nil {
		log.Fatal(err)
	}

	coords := make([]Coord, 0)
	counts := make(map[string]int)
	for i := 0; i+3 < len(bin); i = i + 1 {
		a := int(bin[i])
		b := int(bin[i+1])
		c := int(bin[i+2])
		d := int(bin[i+3])

		x := a | (b << 8)
		y := c | (d << 8)

		var xF float64 = (float64(x) / float64(MAX_DIM)) * float64(IMAGE_WIDTH)
		var yF float64 = (float64(y) / float64(MAX_DIM)) * float64(IMAGE_WIDTH)

		coord := Coord{
			X: int(math.Round(xF)),
			Y: int(math.Round(yF)),
		}
		key := fmt.Sprintf("%d-%d", coord.X, coord.Y)
		log.Println(key)
		keyCount, hasKey := counts[key]
		if !hasKey {
			keyCount = 0
		}
		counts[key] = keyCount + 1
		coords = append(coords, coord)
	}

	maxCount := 0
	for _, v := range counts {
		if maxCount < v {
			maxCount = v
		}
	}

	r := image.Rect(0, 0, int(IMAGE_WIDTH), int(IMAGE_HEIGHT))
	img := image.NewRGBA(r)

	for _, coord := range coords {
		//key := fmt.Sprintf("%d-%d", coord.X, coord.Y)
		img.Set(coord.X, coord.Y, color.RGBA{255, 255, 255, 255})
	}

	f, err := os.Create("img.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println(maxCount)
}
