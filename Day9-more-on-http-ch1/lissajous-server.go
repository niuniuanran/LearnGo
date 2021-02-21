package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", drawLissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func drawLissajous(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	b := r.Form["brightness"]
	if len(b) > 0 {
		brightness, err := strconv.Atoi(r.Form["brightness"][0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintln(w, "Unable to parse brightness!")
			return
		}
		if brightness < 0 || brightness > 254 {
			brightness = 0
		}
		lissajous(w, brightness)
	} else {
		lissajous(w, 0)
	}
}

func lissajous(w io.Writer, brightness int) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	var palette = []color.Color{color.White}
	for i := 0; i < 40; i++ {
		color := color.RGBA{uint8(rand.Intn(255-brightness) + brightness), uint8(rand.Intn(255-brightness) + brightness), uint8(rand.Intn(255-brightness) + brightness), 0xff}
		palette = append(palette, color)
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(math.Abs(t)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay) // The successive delay times, one per frame, in 100ths of a second.
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(w, &anim)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
