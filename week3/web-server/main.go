package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

const port = "localhost:8089"

func main() {
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	lissajous(w)
	// })
	http.HandleFunc("/count", counter)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles := 5
		if s := r.URL.Query().Get("cycles"); s != "" {
			if c, err := strconv.Atoi(s); err != nil {
				log.Printf("Invalid cycles value %q: %v", s, err)
				fmt.Fprintf(w, "Invalid cycles parameter . using default of %d:Error: %v", cycles, err)
			} else {
				cycles = c
			}
		}
		if cycles <= 0 {
			cycles = 5
			fmt.Fprintf(w, "Cycles must be positive. Using default of %d.\n", cycles)
		}
		lissajous(w, cycles)
	})

	println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
func handler(w http.ResponseWriter, r *http.Request) {

	// mu.Lock()
	// count++
	// mu.Unlock()
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "%q %q proto= %s/n", r.Method, r.URL, r.Proto)
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

}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
func lissajous(out io.Writer, cycles int) {
	const (
		// cycles  = 20   // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 200   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	if cycles < 1 {
		cycles = 1
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	if err := gif.EncodeAll(out, &anim); err != nil {
		log.Printf("Error encoding GIF: %v", err) // Log encoding errors
	}
}
