package utils

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

// Prepare takes an array of file paths and converts them to images
func Prepare(paths []string) []*bimg.Image {
	var imgs []*bimg.Image
	for _, path := range paths {
		imgs = append(imgs, NewImage(path))
	}
	return imgs
}

// NewImage Reads a path into an image object
func NewImage(path string) *bimg.Image {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return bimg.NewImage(buffer)
}
