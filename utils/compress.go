package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CompressDir(pathToZip, dstPath string) error {
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	myZip := zip.NewWriter(dstFile)
	err = filepath.Walk(pathToZip, func(filePath string, info os.FileInfo, err error) error {
		// if zip file same name with dir file name to skip it
		if filepath.Base(filePath) == filepath.Base(dstPath) {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		relPath := strings.TrimPrefix(filePath, filepath.Dir(pathToZip))
		zipFile, err := myZip.Create(relPath[1:])
		fmt.Println(pathToZip, filePath, relPath)
		if err != nil {
			return err
		}
		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	err = myZip.Close()
	if err != nil {
		return err
	}
	return nil
}
