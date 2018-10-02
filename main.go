package main

import (
	"bytes"
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
)

var root = flag.String("root", ".", "file system path")

func main() {
	http.HandleFunc("/blue/", blueHandler)
	//http.HandleFunc("/red/", redHandler)
	http.Handle("/", http.FileServer(http.Dir(*root)))
	log.Println("Listening on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen and serve:", err)
	}
}

func blueHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	writeImage(w, &img)
}

func writeImage(w http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)

	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode the image")
	}

	w.Header().Set("Content-Type", "image/jpg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image")
	}

}
