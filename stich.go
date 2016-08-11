package main

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"
)

var (
	zero = image.Point{0, 0}
)

func stitch(images []image.Image) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	for i, simg := range images {
		draw.Draw(img, simg.Bounds().Add(image.Point{(i % 3) * 100, (i / 3) * 100}), simg, zero, draw.Src)
	}
	return img
}

func loadImages(fileNames []string) []image.Image {
	var images []image.Image
	for _, s := range fileNames {
		f, _ := os.OpenFile(s, os.O_RDONLY, 0644)
		img, _ := jpeg.Decode(f)
		images = append(images, img)
	}
	return images
}

func main() {
	fileNames := []string{
		"./resized/ingredients/apples/aplle.jpg",
		"./resized/ingredients/apples/aplle.jpg",
		"./resized/ingredients/apples/aplle.jpg",
		"./resized/ingredients/apples/aplle.jpg",
		"./resized/ingredients/apples/aplle.jpg",
		"./resized/ingredients/apples/aplle.jpg",
		"./resized/ingredients/apples/aplle.jpg",
	}
	images := loadImages(fileNames)
	img := stitch(images)
	b := bytes.NewBuffer(nil)
	jpeg.Encode(b, img, nil)
	ioutil.WriteFile("./new.jpg", b.Bytes(), 0644)
}
