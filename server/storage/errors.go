package storage

import "fmt"

type StorageError struct {
	Cause error
}

func (err *StorageError) Error() string {
	return fmt.Sprintf("storage error: %v", err.Cause)
}
