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
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

type config struct {
	cycles float64
	res float64
	size int
	nframes int
	delay int
}

func main()  {
	//http.HandleFunc("/", handler) // 每個請求呼叫處理程序
	//http.HandleFunc("/count", counter)

	conf := config{
		cycles: 5,
		res: 0.001,
		size: 100,
		nframes: 64,
		delay: 8,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		confs := r.URL.Query()
		for i, c := range confs {
			if i == "cycles" {
				conf.cycles, _ = strconv.ParseFloat(c[0], 64)
			}
		}

		lissajous(w, conf)
	})

	log.Fatal(http.ListenAndServe("localhost:2030", nil))
	return
}

// 處理程序回應所請求 URL 的 Path 元件
// 回應 HTTP 請求的處理程序
func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// 若有兩個並行的請求同時嘗試更新 count，或許無法正確的遞增，會有 race condition 競爭條件的產生
// 因此必須確保每次最多只有一個 goroutine 存取該變數，為 使用 mu.Lock and unLock 圍繞 count 的目的
func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, conf config)  {

	freq := rand.Float64() * 3.0 // y 振盪相對頻率
	anim := gif.GIF{LoopCount: conf.nframes}
	phase := 0.0 // 相位差
	for i := 0; i < conf.nframes; i++ {
		rect := image.Rect(0, 0, 2*conf.size+1, 2*conf.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < conf.cycles*2*math.Pi; t += conf.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(conf.size+int(x*float64(conf.size)+0.5), conf.size+int(y*float64(conf.size)+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, conf.delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}