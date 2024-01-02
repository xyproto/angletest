package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

func angleLength(input uint8) (uint8, uint8) {
	const twopi = 2 * math.Pi
	const steps = 10
	if input == 0 {
		return 0, 0
	}
	length := (input / steps) * steps
	inputFloat64 := float64(input) / 255.0
	angleByte := uint8(math.Mod(inputFloat64*steps*twopi, twopi) * (255.0 / twopi))
	return angleByte, length
}

func drawLine(img draw.Image, start, end image.Point, col color.Color) {
	dx := math.Abs(float64(end.X - start.X))
	dy := math.Abs(float64(end.Y - start.Y))
	sx := -1
	sy := -1
	if start.X < end.X {
		sx = 1
	}
	if start.Y < end.Y {
		sy = 1
	}
	err := dx - dy
	for {
		img.Set(start.X, start.Y, col)
		if start.X == end.X && start.Y == end.Y {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			start.X += sx
		}
		if e2 < dx {
			err += dx
			start.Y += sy
		}
	}
}

func drawCircle(img draw.Image, center image.Point, radius int, col color.Color) {
	for dY := -radius; dY <= radius; dY++ {
		for dX := -radius; dX <= radius; dX++ {
			if dX*dX+dY*dY <= radius*radius {
				img.Set(center.X+dX, center.Y+dY, col)
			}
		}
	}
}

func drawVector(angle, length uint8) *image.RGBA {
	const size = 512
	center := image.Point{size / 2, size / 2}
	radius := size / 2

	img := image.NewRGBA(image.Rect(0, 0, size, size))

	// Draw white background
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Draw circle
	drawCircle(img, center, radius-1, color.Black)

	// Calculate vector end point
	endX := center.X + int(float64(length)*math.Cos(float64(angle)*(2*math.Pi/255)))
	endY := center.Y + int(float64(length)*math.Sin(float64(angle)*(2*math.Pi/255)))

	// Draw vector line
	drawLine(img, center, image.Point{endX, endY}, color.RGBA{255, 0, 0, 255})

	return img
}

func saveImage(img *image.RGBA, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}

func main() {
	var (
		filename string
		angle    byte
		length   byte
		img      *image.RGBA
	)
	for i := 0; i <= 255; i++ {
		filename = fmt.Sprintf("img/vector%03d.png", i)
		angle, length = angleLength(uint8(i))
		fmt.Printf("![input %d](%s)\n\n", i, filename)
		fmt.Printf("    input %d -> angle %f, length %f\n\n", i, float64(angle)/255.0, float64(length)/255.0)
		img = drawVector(angle, length)
		if err := saveImage(img, filename); err != nil {
			fmt.Printf("error saving image: %v\n", err)
		}
	}
}
