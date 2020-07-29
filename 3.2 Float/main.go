package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320     // 畫布大小
	cells = 100                  // 單元格的個數
	xyrange = 30.0               // 坐標軸的範圍
	xyscale = width / 2 / xyrange  // x, y 軸上每個單位長度的像素
	zscale = height * 0.4        // z 軸上每個單位長度的像素
	angle = math.Pi / 6          // x, y 軸的角度 (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30), cos(30)

type zFunc func(x, y float64) float64

/*
os.Stdin 標準輸入
os.Stdout 標準輸出
os.Stderr 標準錯誤輸出
 */

func main()  {

	// Practice 3.4 http
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		createSVG(w)
	})

	log.Fatal(http.ListenAndServe(":2333", nil))

	//claim := "Claim to use: saddle or eggBox or f"
	//if len(os.Args) < 2 {
	//	fmt.Fprintln(os.Stderr, claim)
	//	os.Exit(1)
	//}
	//
	//var f zFunc
	//
	//switch os.Args[1] {
	//
	//case "saddle":
	//	f = saddle
	//case "eggBox":
	//	f = eggBox
	//default:
	//	fmt.Fprintln(os.Stderr, claim)
	//	os.Exit(1)
	//}
	//
	//createSVG(os.Stdout, f)
}

func createSVG(w io.Writer)  {
	zmin, zmax := minmax()
	//fmt.Printf( "<svg xmlns='http://www.w3.org/2000/svg' " +
	//	"style='stroke:grey; fill: white; stroke-width: 0.7' " +
	//	"width='%d' height='%d'>", width, height)

	// Practice 3.3 多邊形上色
	fmt.Fprintf( w, "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke:grey; fill: white; stroke-width: 0.7' " +
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j< cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			// Practice 3.1 函式f 回傳非有限的 float64 值， 略過無效的多邊形
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy){
				continue
			}

			//fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)

			// Practice 3.3 多邊形上色
			fmt.Fprintf(w, "<polygon style='stroke %s; fill: #222222' points='%g,%g %g,%g %g,%g %g,%g'/>\n", color(i, j, zmin, zmax), ax, ay, bx, by, cx, cy, dx, dy)

		}
	}

	fmt.Println("</svg>")
}

//func createSVG(w io.Writer, f zFunc)  {
//	zmin, zmax := minmax()
//	//fmt.Printf( "<svg xmlns='http://www.w3.org/2000/svg' " +
//	//	"style='stroke:grey; fill: white; stroke-width: 0.7' " +
//	//	"width='%d' height='%d'>", width, height)
//
//	// Practice 3.3 多邊形上色
//	fmt.Fprintf( w, "<svg xmlns='http://www.w3.org/2000/svg' " +
//		"style='stroke:grey; fill: white; stroke-width: 0.7' " +
//		"width='%d' height='%d'>", width, height)
//
//	for i := 0; i < cells; i++ {
//		for j := 0; j< cells; j++ {
//			ax, ay := corner(i+1, j, f)
//			bx, by := corner(i, j, f)
//			cx, cy := corner(i, j+1, f)
//			dx, dy := corner(i+1, j+1, f)
//
//			// Practice 3.1 函式f 回傳非有限的 float64 值， 略過無效的多邊形
//			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy){
//				continue
//			}
//
//			//fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
//
//			// Practice 3.3 多邊形上色
//			fmt.Fprintf(w, "<polygon style='stroke %s; fill: #222222' points='%g,%g %g,%g %g,%g %g,%g'/>\n", color(i, j, zmin, zmax), ax, ay, bx, by, cx, cy, dx, dy)
//
//		}
//	}
//
//	fmt.Println("</svg>")
//}

// minmax returns the min and max values for z given the min/max values of x
// and y and assuming a square domain.
func minmax() (min float64, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xoff := 0; xoff <= 1; xoff++ {
				for yoff := 0; yoff <= 1; yoff++ {
					x := xyrange * (float64(i+xoff)/cells - 0.5)
					y := xyrange * (float64(j+yoff)/cells - 0.5)
					z := f(x, y)
					if math.IsNaN(min) || z < min {
						min = z
					}
					if math.IsNaN(max) || z > max {
						max = z
					}
				}
			}
		}
	}
	return
}

func color(i, j int, zmin, zmax float64) string {
	min := math.NaN()
	max := math.NaN()
	for xoff := 0; xoff <= 1; xoff++ {
		for yoff := 0; yoff <= 1; yoff++ {
			x := xyrange * (float64(i+xoff)/cells - 0.5)
			y := xyrange * (float64(j+yoff)/cells - 0.5)
			z := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}

			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	color := ""

	if math.Abs(max) > math.Abs(min) {
		red := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zmax)) * 255
		if red > 255 {
			red = 255
		}

		color = fmt.Sprintf("#%02x0000", int(red))
	} else {
		blue := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zmin)) * 255
		if blue > 255 {
			blue = 255
		}

		color = fmt.Sprintf("#0000%02x", int(blue))
	}

	return color
}

func corner(i, j int) (float64, float64) {
	// 找出(i, j)格的點 (x, y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 計算 z 高度
	z := f(x, y)

	// 投射(x, y, z) 到2-D畫布(sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

//func corner(i, j int, f zFunc) (float64, float64) {
//	// 找出(i, j)格的點 (x, y)
//	x := xyrange * (float64(i)/cells - 0.5)
//	y := xyrange * (float64(j)/cells - 0.5)
//
//	// 計算 z 高度
//	z := f(x, y)
//
//	// 投射(x, y, z) 到2-D畫布(sx, sy)
//	sx := width/2 + (x-y)*cos30*xyscale
//	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
//
//	return sx, sy
//}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // 與(0, 0)的距離

	return math.Sin(r) / r
}

// Practice 3.2 蛋盒
func eggBox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

// Practice 3.2 馬鞍
func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	r := y*y/a2 - x*x/b2

	return r
}
