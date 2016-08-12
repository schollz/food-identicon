package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func resizeEverything() {
	searchDir := "./ingredients/"

	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		// if strings.Contains(path, "olive-oil") {
		fileList = append(fileList, path)
		os.MkdirAll(filepath.Join("resized", path, "../"), os.ModePerm)
		// }
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range fileList {
		if exists(path.Join("resized", f)) {
			continue
		}

		if strings.Contains(f, ".jpg") || strings.Contains(f, ".JPG") {
		} else {
			continue
		}

		file, err := os.Open(f)
		if err != nil {
			continue
		}

		// decode jpeg into image.Image
		img, err := jpeg.Decode(file)
		if err != nil {
			continue
		}
		file.Close()
		fmt.Println(f)

		// resize to width 100 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(100, 100, img, resize.Lanczos3)

		out, err := os.Create(path.Join("resized", f))
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)

	}
}
