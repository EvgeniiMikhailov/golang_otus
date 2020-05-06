package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func openSourceFileWithSeek(fromPath string, offset int64) (*os.File, error) {
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
		_, err = source.Seek(offset, 0)
		if err != nil {
			return nil, err
		}
	}

	return source, nil
}

func calculateActualLimit(source *os.File, limit, offset int64) int64 {
	sourceFI, err := source.Stat()
	if err != nil {
		return 0
	}
	fileSize := sourceFI.Size()
	if limit == 0 || limit > fileSize-offset {
		limit = fileSize - offset
	}
	return limit
}

func Copy(fromPath string, toPath string, offset, limit int64) error {
	source, err := openSourceFileWithSeek(fromPath, offset)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer destination.Close()

	limit = calculateActualLimit(source, limit, offset)

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(source)

	_, err = io.CopyN(destination, barReader, limit)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	bar.Finish()

	return nil
}
