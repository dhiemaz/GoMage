package actions

import (
	"GoMage/internal/image"
	"GoMage/internal/io"
	"fmt"
	"os"
)

func AdjustTemperature(inputPath, outputPath string, temperature float64) {

	fmt.Println(inputPath, outputPath, temperature)
	inputImage, err := io.LoadFile(inputPath)
	if err != nil {
		fmt.Println(fmt.Sprintf("error loading image %s : %v", inputPath, err))
		os.Exit(1)
	}

	outputImage := image.Temperature(inputImage, temperature)

	err = io.SaveFile(outputPath, outputImage)
	if err != nil {
		fmt.Println(fmt.Sprintf("error saving image %s : %v", outputPath, err))
		os.Exit(1)
	}

	fmt.Println("Image temperature adjusted successfully")
}
