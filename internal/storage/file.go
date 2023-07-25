package storage

import (
	"bytes"
	"strconv"
	"strings"
	"time"
)

type File struct {
	buffer *bytes.Buffer
	name   string
}

func NewFile() *File {
	return &File{buffer: &bytes.Buffer{}}
}

func (f *File) Write(chunk []byte) error {
	_, err := f.buffer.Write(chunk)

	return err
}

func (f *File) SetName(clientFilePath string) {
	breadcrumbs := strings.Split(clientFilePath, "/")
	fileName := breadcrumbs[len(breadcrumbs)-1]

	f.name = strconv.Itoa(int(time.Now().Unix())) + "_" + fileName
}

func (f *File) GetName() string {
	return f.name
}
