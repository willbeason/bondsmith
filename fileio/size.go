package fileio

import (
	"fmt"
	"os"
)

func CalculateSize(filepath string) (int64, error) {
	stat, err := os.Stat(filepath)
	if err != nil {
		return 0, fmt.Errorf("statting file: %w", err)
	}

	return stat.Size(), nil
}

func CalculateSizes(filepaths []string) (int64, error) {
	var totalSize int64 = 0
	for _, filepath := range filepaths {
		size, err := CalculateSize(filepath)
		if err != nil {
			return 0, err
		}
		totalSize += size
	}

	return totalSize, nil
}
