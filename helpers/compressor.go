package helpers

import (
	"errors"
	"strings"

	"gopkg.in/h2non/bimg.v1"
)

func CompressToPNG(fileName string) error {
	options := bimg.Options{
		Width:   200,
		Height:  200,
		Quality: 95,
		Type:    bimg.PNG,
	}

	buffer, err := bimg.Read(fileName)
	if err == nil {
		newImage, err := bimg.NewImage(buffer).Process(options)
		if err == nil {
			//save file
			url := fileName[:strings.LastIndexByte(fileName, '.')] + "-compressed"
			bimg.Write(url+".png", newImage)
		} else {
			return errors.New("failed to compress image")
		}
	} else {
		return errors.New("failed to load image")
	}

	return nil
}
