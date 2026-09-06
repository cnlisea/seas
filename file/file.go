package file

import (
	"errors"
	"os"
)

type File struct {
	Path string
	File *os.File
}

func New(path string) *File {
	return &File{
		Path: path,
	}
}

func (f *File) Init() error {
	var err error
	f.File, err = os.OpenFile(f.Path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) WriteString(s string) (int, error) {
	if f.File == nil {
		return 0, errors.New("file not init")
	}

	return f.File.WriteString(s)
}

func (f *File) Close() error {
	if f.File == nil {
		return nil
	}
	return f.File.Close()
}
