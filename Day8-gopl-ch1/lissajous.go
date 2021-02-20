package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

func main() {
	var palette = []color.Color{color.White}
	for i := 0; i < 255; i++ {
		palette = append(palette, color.RGBA{uint8(rand.Intn(100) + 155), uint8(rand.Intn(100) + 155), uint8(rand.Intn(100) + 155), 0xff})
	}
	lissajous(os.Stdout, palette)
}

func lissajous(out io.Writer, palette []color.Color) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
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
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
