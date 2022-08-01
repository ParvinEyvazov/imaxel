package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"image"
	"log"
	"strings"
	"syscall/js"

	_ "image/jpeg"
	_ "image/png"

	"github.com/ParvinEyvazov/imaxel/core"
)

func main() {
	done := make(chan struct{}, 0)

	js.Global().Set("difference", js.FuncOf(difference))

	<-done
}

func difference(this js.Value, args []js.Value) interface{} {

	arr1 := args[0]
	arr2 := args[1]

	buff1 := make([]uint8, arr1.Get("byteLength").Int())
	buff2 := make([]uint8, arr2.Get("byteLength").Int())

	js.CopyBytesToGo(buff1, arr1)
	js.CopyBytesToGo(buff2, arr2)

	image1Base64, err := core.Image{
		Bytes: buff1,
	}.GetBase64()
	mustError(err)

	image2Base64, err := core.Image{
		Bytes: buff2,
	}.GetBase64()
	mustError(err)

	reader1 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image1Base64))
	reader2 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image2Base64))

	image1, _, err := image.Decode(reader1)
	mustError(err)

	image2, _, err := image.Decode(reader2)
	mustError(err)

	bounds1 := image1.Bounds()
	bounds2 := image2.Bounds()

	err = equalPixelCount(bounds1, bounds2)
	mustError(err)

	differentPixel := []image.Point{}

	for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
		for x := bounds1.Min.X; x < bounds1.Max.X; x++ {
			r1, g1, b1, a1 := image1.At(x, y).RGBA()
			r2, g2, b2, a2 := image2.At(x, y).RGBA()

			pixelEquity := r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2

			if !pixelEquity {
				differentPixel = append(differentPixel, image.Point{X: x, Y: y})
			}

		}
	}

	jsonForm, err := json.Marshal(differentPixel)
	mustError(err)

	return string(jsonForm)
}

func equalPixelCount(bounds1, bounds2 image.Rectangle) (err error) {

	equality := bounds1.Min.X == bounds2.Min.X ||
		bounds1.Max.X == bounds2.Max.X ||
		bounds1.Min.Y == bounds2.Min.Y ||
		bounds1.Max.Y == bounds2.Max.Y

	if !equality {

		err = errors.New("Images are not pixely equal")

		return
	}

	return
}

func mustError(err error) {

	if err != nil {
		log.Fatalln(err)
	}
}
