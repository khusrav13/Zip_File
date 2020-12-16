package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func archive(file, aim string) error {
	value, err := os.Create(aim)
	if err != nil {
		return err
	}
	defer value.Close()

	archiveit := zip.NewWriter(value)
	defer archiveit.Close()

	info, err := os.Stat(file)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(file)
	}

	filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, file))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archiveit.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

func main()  {
	archive("C:/Users/KHUSRAV/go/src/ZipArchive/index.html", "C:/Users/KHUSRAV/go/src/ZipArchive/indexx.zip")
}