package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func openSourceFile(fromPath string, offset, limit int64) (*os.File, int64, error) {
	readLimit := limit

	source, err := os.Open(fromPath)
	if err != nil {
		return nil, readLimit, err
	}
	sourceFI, err := source.Stat()
	if err != nil {
		return nil, readLimit, err
	}

	if !sourceFI.Mode().IsRegular() {
		return nil, readLimit, ErrUnsupportedFile
	}

	fileSize := sourceFI.Size()
	if offset > fileSize {
		return nil, readLimit, ErrOffsetExceedsFileSize
	}

	if offset != 0 {
		source.Seek(offset, 0)
	}

	if limit == 0 || limit > fileSize-offset {
		readLimit = fileSize - offset
	}

	return source, readLimit, nil
}

func Copy(fromPath string, toPath string, offset, limit int64) error {
	source, limit, err := openSourceFile(fromPath, offset, limit)
	defer source.Close()
	if err != nil {
		return err
	}

	destination, err := os.Create(toPath)
	defer destination.Close()
	if err != nil {
		return err
	}

	_, err = io.CopyN(destination, source, limit)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}
