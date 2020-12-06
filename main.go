package main

import (
	"fmt"
	"os"
	"path/filepath"

	"locals/imgprocess"
	"locals/utils"
)

const dataPath = "./data"

func scanDataPath(rootPath string) []string {
	var files []string

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	return files[1:]
}

// Applies some function to an array of images
// Can be improved to use concurrency
// func execute(a Action) {
// }

func main() {
	files := scanDataPath(dataPath)
	imgs := utils.Prepare(files)
	fmt.Println(files)

	for _, img := range imgs {
		fmt.Println(imgprocess.Size(img))
	}
}
