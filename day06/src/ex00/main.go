package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))

	colors := []color.Color{
		color.RGBA{R: 0, G: 0, B: 0, A: 0},       // Background
		color.RGBA{R: 255, G: 0, B: 128, A: 255}, // Magenta
		color.RGBA{R: 128, G: 0, B: 255, A: 255}, // Purple
		color.RGBA{R: 0, G: 0, B: 0, A: 255},     // Black
	}

	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			colorX := float64(x - 150)
			colorY := float64(y - 150)

			idx := 0
			
			if colorX >= 0 && colorY >= 0 && colorX+colorY <= 150 {
				idx = int((colorX+colorY)/10) % len(colors)
			} 
			
			if colorX <= 0 && colorY >= 0 && -colorX+colorY <= 150 {
				idx = int((-colorX+colorY)/10) % len(colors)
			} 
			
			if colorX <= 0 && colorY <= 0 && -colorX-colorY <= 150 {
				idx = int((-colorX-colorY)/10) % len(colors)
			} 

			if colorX >= 0 && colorY <= 0 && colorX-colorY <= 150 {
				idx = int((colorX-colorY)/10) % len(colors)
			}

			img.Set(x, y, colors[idx])
		}
	}

	file, err := os.Create("amazing_logo.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
}
