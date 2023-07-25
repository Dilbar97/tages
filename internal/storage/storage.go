package storage

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Storage struct {
	Dir string
}

func New(dir string) Storage {
	return Storage{Dir: dir}
}

func (s Storage) Store(file *File) error {
	if _, err := os.Stat(s.Dir); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(s.Dir, 0777); err != nil {
			log.Println(err)
			return err
		}
	}

	if err := os.WriteFile(s.Dir+file.name, file.buffer.Bytes(), 0777); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
