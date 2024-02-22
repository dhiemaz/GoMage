package actions

import (
	"github.com/fogleman/gg"
	"image/color"
)

func ProcessImage(inputPath, outputPath string, temperature uint32) error {
	// Open the image file
	img, err := gg.LoadImage(inputPath)
	if err != nil {
		return err
	}

	// Adjust the temperature of the image
	adjustedImg := adjustTemperature(gg.NewContextForImage(img), temperature)

	// Save the adjusted image to the output path
	err = adjustedImg.SavePNG(outputPath)
	if err != nil {
		return err
	}

	return nil
}

// adjustTemperature applies color tone adjustment to simulate warmer or cooler temperatures.
func adjustTemperature(img *gg.Context, temperature uint32) *gg.Context {
	width, height := img.Width(), img.Height()
	result := gg.NewContext(width, height)

	result.DrawImage(img.Image(), 0, 0)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.Image().At(x, y).RGBA()

			// Adjust temperature
			newR := r + temperature
			newG := g + temperature
			newB := b + temperature

			// Apply bounds to prevent overflow
			newR = clamp(newR, 0, 65535)
			newG = clamp(newG, 0, 65535)
			newB = clamp(newB, 0, 65535)

			// Create a new color with adjusted RGB components
			newColor := color.RGBA{
				R: uint8(newR >> 8),
				G: uint8(newG >> 8),
				B: uint8(newB >> 8),
				A: 255,
			}

			result.SetColor(newColor)
			result.SetPixel(x, y)
		}
	}

	return result
}

// clamp ensures that a value stays within a specified range.
func clamp(value, min, max uint32) uint32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
