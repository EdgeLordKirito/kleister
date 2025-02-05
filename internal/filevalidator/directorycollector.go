package filevalidator

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))
var last string

func CollectImageFiles(directories []string) ([]string, int64, error) {
	var imageFiles []string
	timestamp := time.Now().Unix()

	for _, dir := range directories {
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, 0, err
		}

		for _, entry := range entries {
			if !entry.IsDir() && IsImageFile(entry.Name()) {
				imageFiles = append(imageFiles, filepath.Join(dir, entry.Name()))
			}
		}
	}

	return imageFiles, timestamp, nil
}

func PickRandomFile(files []string) (string, error) {
	if len(files) == 0 {
		return "", fmt.Errorf("No Files in directories") // No image files found
	}
	if len(files) == 1 {
		return files[0], fmt.Errorf("Singular File found in directories")
	}
	file := files[rng.Intn(len(files))]
	if file == last {
		//rerun until new file is selected
		for i := 0; i < 5; i++ {
			file = files[rng.Intn(len(files))]
			if file != last {
				break
			}
		}
	}
	last = file
	return file, nil
}
