package utils

import (
	"os"
	"path"
)

func createDirIfNonExisting(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

// getBaseFolderPath returns the path to the .changeset folder
func getBaseFolderPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Get the base folder path
	baseFolderPath := path.Join(wd, ".changeset")
	err = createDirIfNonExisting(baseFolderPath)
	if err != nil {
		return "", nil
	}

	return baseFolderPath, nil
}
