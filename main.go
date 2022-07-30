package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"strings"

	_ "image/png"

	"github.com/ParvinEyvazov/imaxel/core"
)

func main() {

	image1Base64 := core.GetBase64("1.png")
	image2Base64 := core.GetBase64("2.png")

	reader1 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image1Base64))
	reader2 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image2Base64))

	image1, _, err1 := image.Decode(reader1)
	image2, _, err2 := image.Decode(reader2)
	if err1 != nil {
		log.Fatal(err1)
	}
	if err2 != nil {
		log.Fatal(err2)
	}

	bounds1 := image1.Bounds()
	bounds2 := image2.Bounds()

	pixelCountEquity := bounds1.Min.X == bounds2.Min.X ||
		bounds1.Max.X == bounds2.Max.X ||
		bounds1.Min.Y == bounds2.Min.Y ||
		bounds1.Max.Y == bounds2.Max.Y

	if !pixelCountEquity {
		log.Fatal("Images are not pixely equal.")
	}

	for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
		for x := bounds1.Min.X; x < bounds1.Max.X; x++ {
			r1, g1, b1, a1 := image1.At(x, y).RGBA()
			r2, g2, b2, a2 := image2.At(x, y).RGBA()

			pixelEquity := r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2

			if !pixelEquity {
				fmt.Println("different pixel: ", x, "x", y)
			}

		}
	}

}
