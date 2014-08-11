// Package rand provides helper functions to create random ids.
package rand

import (
	"crypto/rand"
	"errors"
)

// exactly 64 characters
const characters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-+"

// Id generates a random id containing only alphanumeric characters.
func Id(length int) (string, error) {

	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	buff := make([]byte, length)
	for i := 0; i < length; i++ {
		z := b[i] >> 2 // z is between 0 and 63
		buff[i] = characters[z]
	}
	return string(buff), nil
}

// IdPrefix generates a random id of the desired length, starting with the given prefix
func IdPrefix(prefix string, length int) (string, error) {
	if len(prefix) > length {
		return "", errors.New("Prefix length should not exceed the required length")
	}

	suffix, err := Id(length - len(prefix))
	if err != nil {
		return "", err
	}

	return prefix + suffix, nil
}
