package actions

import (
	"GoMage/internal/image"
	"GoMage/internal/io"
	"fmt"
)

func AdjustTemperature(inputPath, outputPath string, temperature float64) {
	inputImage, err := io.LoadFile(inputPath)
	if err != nil {
		fmt.Errorf("error loading image: %v", err)
	}

	outputImage := image.AdjustTemperature(inputImage, temperature)

	err = io.SaveFile(outputPath, outputImage)
	if err != nil {
		fmt.Errorf("error saving image: %v", err)
	}

	fmt.Println("Image temperature adjusted successfully")
}
