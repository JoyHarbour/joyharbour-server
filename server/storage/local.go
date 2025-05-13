package storage

import (
	"io"
	"os"
	"path"
)

type LocalStorage struct {
	Path string
}

func (storage *LocalStorage) Retrieve(id string, w io.Writer) error {
	file, err := os.Open(path.Join(storage.Path, id))
	if err != nil {
		return &StorageError{Cause: err}
	}

	_, err = io.Copy(w, file)
	return err
}

func (storage *LocalStorage) Store(id string, r io.Reader) error {
	file, err := os.Create(path.Join(storage.Path, id))
	if err != nil {
		return &StorageError{Cause: err}
	}

	_, err = io.Copy(file, r)
	return err
}
