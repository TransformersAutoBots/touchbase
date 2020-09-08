package logger

import (
    "go.uber.org/zap"
)

const (
    touchBaseErrorKey = "touchBaseError"
)

// convertToField takes a key and an arbitrary value and chooses the best way
// to represent them as a field, failing to reflection-based approach only if
// necessary.
func convertToField(key string, v interface{}) zap.Field {
    return zap.Any(key, v)
}

func Attribute(key string, v interface{}) zap.Field {
    return convertToField(key, v)
}

func TouchBaseError(v interface{}) zap.Field {
    return convertToField(touchBaseErrorKey, v)
}
