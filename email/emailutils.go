package email

import (
    "crypto/rand"

    "github.com/autobots/touchbase/constants"
)

func chunkSplit(body string, limit int, end string) string {
    var charSlice []rune
    for _, char := range body {
        charSlice = append(charSlice, char)
    }

    var result string
    for len(charSlice) >= 1 {
        result = result + string(charSlice[:limit]) + end
        charSlice = charSlice[limit:]
        if len(charSlice) < limit {
            limit = len(charSlice)
        }
    }
    return result
}

func randomString(stringSize int, randomType string) string {
    var dictionary string
    switch randomType {
    case constants.AlphaNumericType:
        dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    case constants.AlphaType:
        dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    case constants.NumericType:
        dictionary = "0123456789"
    }

    var bytes = make([]byte, stringSize)
    _, err := rand.Read(bytes)
    if err != nil {
        return dictionary[:stringSize]
    }
    for k, v := range bytes {
        bytes[k] = dictionary[v%byte(len(dictionary))]
    }
    return string(bytes)
}
