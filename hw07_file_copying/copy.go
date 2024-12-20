package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromInfo, err := getFileInfo(fromPath)
	if err != nil {
		return err
	}

	if fromInfo.Size() == 0 {
		return ErrUnsupportedFile
	}

	if offset > fromInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	src, err := openFile(fromPath)
	if err != nil {
		return err
	}
	defer func() {
		errC := src.Close()
		if errC != nil {
			log.Println(errC)
		}
	}()

	err = fileOffset(src, offset)
	if err != nil {
		return err
	}

	copyLength, err := getCopyLength(fromInfo.Size(), offset, limit)
	if err != nil {
		return err
	}

	dst, err := fileCreate(toPath)
	if err != nil {
		return err
	}

	defer func() {
		errC := dst.Close()
		if errC != nil {
			log.Println(errC)
		}
	}()

	bar := pb.Full.Start64(copyLength)
	barReader := bar.NewProxyReader(src)
	_, err = io.CopyN(dst, barReader, copyLength)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	bar.Finish()

	return nil
}

func getFileInfo(filePath string) (os.FileInfo, error) {
	if len(filePath) == 0 {
		return nil, fmt.Errorf("empty file path")
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	return fileInfo, nil
}

func openFile(filePath string) (*os.File, error) {
	if len(filePath) == 0 {
		return nil, fmt.Errorf("empty file path")
	}

	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func fileOffset(file *os.File, offset int64) error {
	if offset < 0 {
		return fmt.Errorf("offset parameter cannot be less than zero")
	}

	_, err := file.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	return nil
}

func getCopyLength(fileSize, offset, limit int64) (int64, error) {
	if limit < 0 {
		return 0, fmt.Errorf("limit cannot be leass than zero")
	}

	var copyLength int64

	copyLength = fileSize - offset

	if limit > 0 {
		copyLength = min(copyLength, limit)
	}

	return copyLength, nil
}

func fileCreate(filePath string) (*os.File, error) {
	if len(filePath) == 0 {
		return nil, fmt.Errorf("empty file path")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}
