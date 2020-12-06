package utils

import (
    "encoding/json"
    "os"
    "path/filepath"

    "github.com/autobots/touchbase/constants"
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

// IsNotEmptyString checks if string is not empty.
//
// Args:
//   value: the string value
// Return:
//   true: if string is not empty
//   false: otherwise
func IsNotEmptyString(value string) bool {
    return !IsEmptyString(value)
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

// PrettyJson is like Marshal json but applies Indent to format the output.
// Each JSON element in the output will begin on a new line beginning with prefix
// followed by one or more copies of indent according to the indentation nesting.
//
// Args:
//   v: the struct data
// Return:
//   the indented json data
//   error: if failed to marshal and indent json data
func PrettyJson(v interface{}) ([]byte, error) {
    bytes, err := json.MarshalIndent(v, constants.JsonPrefix, constants.JsonIntend)
    if err != nil {
        return nil, err
    }
    return bytes, nil
}

// CleanFilePath returns the shortest path name equivalent to path by purely
// lexical processing. It applies the following rules iteratively until no
// further processing can be done.
//
// Args:
//   v: the dir/file path
// Return:
//   the cleaned dir/file path
func CleanFilePath(path string) string {
    if IsEmptyString(path) {
        return path
    }

    absPath, err := filepath.Abs(path)
    if err != nil {

    }
    return filepath.Clean(absPath)
}

func GetAbsPath(path string) (string, error) {
    if IsEmptyString(path) {
        return path, nil
    }
    absPath, err := filepath.Abs(path)
    if err != nil {
        return "", err
    }
    return CleanFilePath(absPath), nil
}
