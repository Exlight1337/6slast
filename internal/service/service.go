package service

import (
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {

	hasLetters := false

	data = strings.TrimSpace(data)

	if data == "" {
		return "", nil
	}

	for _, r := range data {
		if unicode.IsLetter(r) {
			hasLetters = true
			break
		}
	}

	if hasLetters {
		return morse.ToMorse(data), nil
	}

	return morse.ToText(data), nil
}
