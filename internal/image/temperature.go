package image

import (
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/color"
)

// Temperature function to adjust the temperature of an image
func Temperature(img image.Image, temperature float64) image.Image {
	bounds := img.Bounds()
	newImage := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			originalColor, _ := colorful.MakeColor(img.At(x, y))
			adjustedColor := adjustColor(originalColor, temperature)
			newImage.Set(x, y, adjustedColor)
		}
	}

	return newImage
}

// adjustColor (private) function to adjust the color of a pixel
func adjustColor(c color.Color, temperature float64) color.Color {

	// Extract the RGB values from c
	r, g, b, a := c.RGBA()

	// Convert temperature to a factor between 0.0 and 1.0
	factor := float64(temperature) / 1000.0

	// Adjust the red component to make the color warmer
	r = uint32(float64(r) * (1.0 + factor))

	// Ensure the values are within the valid range (0-65535)
	r = clamp(r, 0, 65535)
	g = clamp(g, 0, 65535)
	b = clamp(b, 0, 65535)

	// Create a new color.RGBA
	adjustedColor := color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}

	return adjustedColor
}

// clamp (private) function to ensure a value is within a specified range
func clamp(value, min, max uint32) uint32 {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}
