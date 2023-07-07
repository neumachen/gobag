package gobag

import (
	"io/ioutil"
	"path/filepath"
)

// CreateTempFile ...
func CreateTempFile(fileContent []byte, directory, dirPrefix, fileName string) (string, string, error) {
	dir, err := ioutil.TempDir(dirPrefix, directory)
	if err != nil {
		return "", "", Errx(err)
	}

	tmpfn := filepath.Join(dir, fileName)
	err = ioutil.WriteFile(tmpfn, fileContent, 0666)
	if err != nil {
		return "", "", Errx(err)
	}
	return dir, tmpfn, nil
}
