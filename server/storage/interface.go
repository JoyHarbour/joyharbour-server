package storage

import "io"

type Storage interface {
	Retrieve(id string, w io.Writer) error
	Store(id string, r io.Reader) error
}
