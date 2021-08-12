package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v2"

	"github.com/desmos-labs/changeset/types"
)

// GetEntriesFolderPath returns the path to the entries folder, or an error if something goes wrong
func GetEntriesFolderPath() (string, error) {
	baseFolderPath, err := getBaseFolderPath()
	if err != nil {
		return "", err
	}

	entriesFolderPath := path.Join(baseFolderPath, "entries")
	err = createDirIfNonExisting(entriesFolderPath)
	if err != nil {
		return "", nil
	}

	return entriesFolderPath, nil
}

// WriteEntry writes the given entry inside a newly created file
func WriteEntry(entry *types.Entry) error {
	entriesFolderPath, err := GetEntriesFolderPath()
	if err != nil {
		return err
	}

	bz, err := yaml.Marshal(entry)
	if err != nil {
		return fmt.Errorf("error while serializing changeset entry: %s", err)
	}

	shasum := sha256.Sum256(bz)
	entryName := fmt.Sprintf("%s.yaml", hex.EncodeToString(shasum[:]))
	entryFilePath := path.Join(entriesFolderPath, entryName)

	err = ioutil.WriteFile(entryFilePath, bz, 0666)
	if err != nil {
		return fmt.Errorf("error while writing entry file %s", err)
	}

	return nil
}

// ReadEntry reads the file at the given path and parses the contents as an Entry object
func ReadEntry(path string) (*types.Entry, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading entry file: %s", err)
	}

	var entry types.Entry
	err = yaml.Unmarshal(bytes, &entry)
	if err != nil {
		return nil, fmt.Errorf("error while parsing entry: %s", err)
	}
	return &entry, nil
}

// GetEntries returns the current list of changeset entries
func GetEntries() ([]*types.Entry, error) {
	entriesFolder, err := GetEntriesFolderPath()
	if err != nil {
		return nil, err
	}

	items, err := ioutil.ReadDir(entriesFolder)
	if err != nil {
		return nil, fmt.Errorf("error while reading the entries folder: %s", err)
	}

	var entries = make([]*types.Entry, len(items))
	for index, item := range items {
		entryFilePath := path.Join(entriesFolder, item.Name())
		entry, err := ReadEntry(entryFilePath)
		if err != nil {
			return nil, err
		}

		entries[index] = entry
	}

	return entries, nil
}
