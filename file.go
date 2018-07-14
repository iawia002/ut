package ut

import (
	"os"
)

// FileSize returns the file size of the specified path file.
func FileSize(filePath string) (int64, bool) {
	file, err := os.Stat(filePath)
	if err != nil {
		return 0, false
	}
	return file.Size(), true
}
