package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand" // Needed for rand.Float64()
	"os"
	"time" // Needed for rand.Seed
)

// palette defines the set of colors available for the GIF.
// Exercise 1.5: The first non-black color is green.
// Exercise 1.6: Additional colors (blue, red, yellow) are added to the palette.
var palette = []color.Color{
	color.Black,                        // 0: Black background
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // 1: Vibrant Green
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // 2: Bright Blue
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // 3: Pure Red
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // 4: Yellow
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, // 5: Magenta
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, // 6: Cyan
}

const (
	cycles  = 5     // Number of complete x oscillator revolutions
	res     = 0.001 // Angular resolution (smaller value makes smoother curves)
	size    = 100   // Image canvas covers [-size..+size]
	nframes = 64    // Number of animation frames
	delay   = 8     // Delay between frames in 10ms units
)

func main() {
	// Seed the random number generator using the current time.
	// This ensures that each run of the program produces a different,
	// unique Lissajous figure (due to the `freq` variable).
	rand.Seed(time.Now().UnixNano())

	// Call the lissajous function to generate the GIF and write it to standard output.
	// To view the GIF, you would typically redirect the output to a file:
	// go run lissajous.go > lissajous.gif
	lissajous(os.Stdout)
}

// lissajous generates a GIF of a Lissajous figure and writes it to the provided io.Writer.
func lissajous(out io.Writer) {
	// freq is the relative frequency of the y oscillator.
	// A random value is used to create diverse patterns.
	freq := rand.Float64() * 3.0

	// anim holds the animated GIF structure.
	// LoopCount:nframes makes the GIF loop continuously.
	anim := gif.GIF{LoopCount: nframes}

	// phase is the phase difference between the x and y oscillators.
	// It changes slightly with each frame to create the animation.
	phase := 0.0

	// numDrawingColors is the number of colors available for drawing the curve,
	// excluding the black background.
	numDrawingColors := len(palette) - 1

	// Loop through each frame of the animation.
	for i := 0; i < nframes; i++ {
		// Create a new image with the specified dimensions (2*size+1 for center-aligned drawing).
		// The image uses the defined palette.
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// Loop to draw the Lissajous curve for the current frame.
		// t represents the angle for the x oscillator.
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// Calculate x and y coordinates based on sine waves.
			// The y-coordinate's frequency and phase are varied.
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			// Exercise 1.6: Calculate the color index in an "interesting way".
			// This method makes the color of the curve change progressively along its path.
			// It normalizes 't' (which represents the angle/progress along the curve)
			// to a value between 0 and 1, then scales it to select a color from the
			// available drawing colors (excluding black).
			colorIndex := uint8(math.Floor(t/(cycles*2*math.Pi)*float64(numDrawingColors))) + 1

			// Ensure the calculated colorIndex is within the valid range of the palette
			// (from 1 to numDrawingColors). This handles potential floating-point
			// inaccuracies or edge cases, preventing out-of-bounds errors.
			if colorIndex >= uint8(len(palette)) {
				colorIndex = uint8(len(palette) - 1) // If too high, use the last color
			} else if colorIndex == 0 {
				colorIndex = 1 // Ensure it's never black for the curve itself
			}

			// Set the color of the pixel at (x, y) on the image.
			// Coordinates are scaled and shifted to fit the image dimensions.
			// The chosen colorIndex determines the color from the 'palette'.
			img.SetColorIndex(int(x*size+size), int(y*size+size), colorIndex)
		}

		// Increment the phase for the next frame, creating the animation.
		phase += 0.1

		// Append the current frame's delay and image to the GIF animation.
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// Encode the entire GIF animation to the provided io.Writer.
	// Errors during encoding are ignored in this example for simplicity.
	gif.EncodeAll(out, &anim)
}
