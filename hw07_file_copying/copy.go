package main

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

const bufSize = 1024

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	wFile, err := os.OpenFile(toPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o644)
	defer closeFile(wFile)
	if err != nil {
		return err
	}
	rFile, err := os.Open(fromPath)
	defer closeFile(rFile)
	if err != nil {
		return err
	}
	stat, err := rFile.Stat()
	if err != nil {
		return err
	}
	size := stat.Size()
	if err := isValidFileSize(rFile, size); err != nil {
		return err
	}
	if offset > size {
		return ErrOffsetExceedsFileSize
	}
	if _, err = rFile.Seek(offset, io.SeekStart); err != nil {
		return err
	}
	reader := setLimit(rFile, limit)

	buff := make([]byte, bufSize)
	bar := pb.Start64(size)
	for {
		n, err := reader.Read(buff)
		bar.Add(n)
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		_, err = wFile.Write(buff[:n])
		if err != nil {
			return err
		}
	}
}

func setLimit(f *os.File, l int64) io.Reader {
	if l > 0 {
		return io.LimitReader(f, l)
	}
	return f
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func isValidFileSize(f *os.File, size int64) error {
	if size == 0 {
		_, err := f.Read(make([]byte, 1))
		if !errors.Is(err, io.EOF) {
			return ErrUnsupportedFile
		}
	}
	return nil
}
