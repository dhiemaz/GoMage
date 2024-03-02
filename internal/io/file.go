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
	if err := imageExtensionValidation(fileName); err != nil {
		return nil, err
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

	// Determine the image format (must be either jpg or jpeg format
	if err := imageExtensionValidation(fileName); err != nil {
		return fmt.Errorf("unsupported file format")
	}

	return jpeg.Encode(file, img, nil)
}

func imageExtensionValidation(fileName string) error {
	if fileName[len(fileName)-3:] != "jpg" && fileName[len(fileName)-4:] != "jpeg" {
		return errors.New("unsupported file format")
	}

	return nil
}
