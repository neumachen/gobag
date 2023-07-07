package gobag

import (
	"os"
	"strings"
)

// CreateTempFile ...
func CreateTempFile(fileContent []byte, dirPrefix, fileName string) (*os.File, error) {
	tmpFile, err := os.Create(
		strings.Join([]string{
			os.TempDir(),
			dirPrefix,
			fileName,
		}, "/"),
	)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(tmpFile.Name(), fileContent, 0666)
	if err != nil {
		return nil, err
	}
	return tmpFile, nil
}
