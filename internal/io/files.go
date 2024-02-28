package io

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

// LoadFile function to load an image from a file
func LoadFile(fileName string) (image.Image, error) {

	if fileName[len(fileName)-3:] != "jpg" || fileName[len(fileName)-4:] != "jpeg" {
		return nil, errors.New("unsupported file format")
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

// SaveFile function to save an image to a file
func SaveFile(fileName string, img image.Image) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	// Determine the image format and encode accordingly
	if fileName[len(fileName)-3:] == "jpg" || fileName[len(fileName)-4:] == "jpeg" {
		return jpeg.Encode(file, img, nil)
	}

	return fmt.Errorf("unsupported file format")
}
