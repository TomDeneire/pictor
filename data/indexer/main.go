package main

import (
	"bufio"
	"os"
	"path/filepath"

	lib "tomdeneire.github.io/pictor/lib"
)

const manifestsDir = "corpus"

func handleFile(path string) error {
	file, openErr := os.Open(path)
	if openErr != nil {
		return openErr
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanLines)

	var manifests []string

	for scanner.Scan() {
		manifests = append(manifests, scanner.Text())
	}

	// Error handling
	scanErr := scanner.Err()
	if scanErr != nil {
		return scanErr
	}
	err := lib.ManifestWrapper(manifests)
	if err != nil {
		return err
	}
	return nil
}

func handleDir(path string, d os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return nil
	}

	err = handleFile(path)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	err := filepath.Walk(manifestsDir, handleDir)
	if err != nil {
		panic(err)
	}
}
