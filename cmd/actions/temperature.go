package actions

import (
	"GoMage/internal/image"
	"GoMage/internal/io"
	"fmt"
	"log"
	"os"
)

func AdjustTemperature(inputPath, outputPath string, temperature float64) {

	inputImage, err := io.LoadFile(inputPath)
	if err != nil {
		log.Fatalf(fmt.Sprintf("error loading image %s : %v", inputPath, err))
		os.Exit(1)
	}

	log.Println(fmt.Sprintf("adjusting image temperature file %s with temperature value %.1f", inputPath, temperature))

	outputImage := image.Temperature(inputImage, temperature)

	err = io.SaveFile(outputPath, outputImage)
	if err != nil {
		log.Fatalf(fmt.Sprintf("error saving image %s : %v", outputPath, err))
		os.Exit(1)
	}

	log.Println(fmt.Sprintf("stored result image to %s", outputPath))

	log.Println("image temperature adjusted successfully")
}
