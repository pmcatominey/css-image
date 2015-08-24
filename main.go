package main

import (
	"flag"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"mime/multipart"
	"net/http"
)

var (
	listenAddr = flag.String("listen", ":8080", "http listen address")
)

type pixel struct {
	R, G, B, A uint8
}

func main() {
	flag.Parse()

	log.Println("listening on", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, http.HandlerFunc(rootHandler)))
}

func processFile(file multipart.File, blockSize int) (pixels [][]pixel, bounds image.Rectangle, err error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return
	}

	bounds = img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y += blockSize {
		row := []pixel{}
		for x := bounds.Min.X; x < bounds.Max.X; x += blockSize {
			r, g, b, a := img.At(x, y).RGBA()
			col := pixel{
				uint8(r % 256),
				uint8(g % 256),
				uint8(b % 256),
				uint8(a % 256),
			}

			row = append(row, col)

		}
		pixels = append(pixels, row)
	}

	return
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var data mainTplData

	if r.Method != "POST" {
		// return page without image
		mainTpl.Execute(w, data)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		panic(err)
	}

	pixels, bounds, err := processFile(file, 1)
	if err != nil {
		panic(err)
	}

	data.Width = bounds.Max.X
	data.Height = bounds.Max.Y
	data.Pixels = pixels
	data.PixelCount = data.Width * data.Height

	mainTpl.Execute(w, data)
}
