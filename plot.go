package plot

import (
	"bufio"
	"image"
	"image/color"
	"image/gif"
	"io"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

type Plot struct {
	xMin, xMax, yMin, yMax, dx, yx float64
	nPixels                        int
}

func NewPlot(xMin, xMax, yMin, yMax float64, nPixels int) *Plot {
	dx := (xMax - xMin) / float64(nPixels)
	yx := (yMax - yMin) / float64(nPixels)
	return &Plot{xMin, xMax, yMin, yMax, dx, yx, nPixels}
}

func (plot *Plot) Dx() float64 {
	return (plot.xMax - plot.xMin) / float64(plot.nPixels)
}

func (plot *Plot) Yx() float64 {
	return (plot.yMax - plot.yMin) / float64(plot.nPixels)
}

func (plot *Plot) XPixel(x float64) int {
	return int((x - plot.xMin) / plot.dx)
}

func (plot *Plot) YPixel(y float64) int {
	return plot.nPixels - int((y-plot.yMin)/plot.yx)
}

func (plot *Plot) Draw(out io.Writer, equations ...func(float64) float64) {
	img := plot.DrawInner(0, equations...)
	gif.EncodeAll(out, &gif.GIF{
		Image: []*image.Paletted{img},
		Delay: []int{0},
	})
}

func (plot *Plot) DrawInner(shift float64, equations ...func(float64) float64) *image.Paletted {
	rect := image.Rect(0, 0, plot.nPixels, plot.nPixels)
	img := image.NewPaletted(rect, palette)
	dx := plot.dx
	for x := plot.xMin; x < plot.xMax; x += dx {
		for _, equation := range equations {
			y := equation(x + shift)
			img.SetColorIndex(plot.XPixel(x), plot.YPixel(y), blackIndex)
		}
	}
	return img
}

func (plot *Plot) DrawMoving(out io.Writer, equations ...func(float64) float64) {
	const (
		cycles  = 5  // number of complete x oscillator revolutions
		nframes = 64 // number of animation frames
		delay   = 8  // delay between frames in 10ms units
	)
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		shift := float64(cycles) * (float64(i) / float64(nframes))
		img := plot.DrawInner(shift, equations...)
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func WriteToFile(name string, generator func(io.Writer)) {
	f, _ := os.Create(name)
	defer f.Close()
	w := bufio.NewWriter(f)
	generator(w)
	w.Flush()
}
