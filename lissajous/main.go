package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0x22, 0x8B, 0x22, 0xff}, color.RGBA{0x00, 0x8B, 0x8B, 0xff}, color.RGBA{0xff, 0x00, 0xff, 0xff}}

const (
	whiteIndex = 0 // 調色盤的第一個顏色
	blackIndex = 1 // 調色盤的下一個顏色
	greenIndex = 2
	cyanIndex = 3
	magentaIndex = 4
)

func main()  {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer)  {
	const (
		cycles = 5 // x 振盪的旋轉數
		res = 0.001 // 角解析度
		size = 100 // 畫布大小
		nframes = 64 // 禎數
		delay = 8 // 以10ms 為單位的禎間隔
	)

	freq := rand.Float64() * 3.0 // y 振盪相對頻率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 相位差
	for i, c := 0, 1; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(c))
		}

		c ++ //修改顏色選擇
		if c >= len(palette) {
			c = 1
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
