package which

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrFileNotFound = errors.New("file not found")
)

func Execute(cmd string) error {
	paths := filepath.SplitList(os.Getenv("PATH"))

	for _, directory := range paths {
		candidateFilePath := filepath.Join(directory, cmd)

		fileInfo, err := os.Stat(candidateFilePath)
		if err == nil {
			fileMode := fileInfo.Mode()

			if fileMode.IsRegular() && fileMode&0111 != 0 {
				fmt.Println(candidateFilePath)
				return nil
			}
		}
	}

	return ErrFileNotFound
}
