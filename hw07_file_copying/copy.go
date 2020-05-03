package main

import (
	"errors"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func openSourceFile(fromPath string, offset int64) (*os.File, error) {
	source, err := os.Open(fromPath)
	if err != nil {
		return nil, err
	}
	sourceFI, err := source.Stat()
	if err != nil {
		return nil, err
	}

	if !sourceFI.Mode().IsRegular() {
		return nil, ErrUnsupportedFile
	}

	if offset > sourceFI.Size() {
		return nil, ErrOffsetExceedsFileSize
	}

	if offset != 0 {
		source.Seek(offset, 0)
	}
	return source, nil
}

func Copy(fromPath string, toPath string, offset, limit int64) error {
	source, err := openSourceFile(fromPath, offset)
	defer source.Close()
	if err != nil {
		return err
	}

	destination, err := os.Create(toPath)
	defer destination.Close()
	if err != nil {
		return err
	}

	return nil
}
