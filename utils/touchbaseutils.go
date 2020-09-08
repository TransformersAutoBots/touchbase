package utils

import (
    "os"
)

// IsEmptyString checks if string is empty.
//
// Args:
//   value: the string value
// Return:
//   true: if string is empty
//   false: otherwise
func IsEmptyString(value string) bool {
    return value == ""
}

// Mkdir uses os.MkdirAll to create a directory named path, along with any
// necessary parents, and return nil, or else returns an error.
//
// Args:
//   path: the dir path
//   perm: the permission for the dir
// Return:
//   error: if failed to to create dir
func Mkdir(path string, perm os.FileMode) error {
    if err := os.MkdirAll(path, perm); err != nil {
        return err
    }
    return nil
}
