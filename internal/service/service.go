package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("empty string")
	}
	if strings.ContainsAny(s, "-.") {
		return morse.ToText(s), nil
	} else {
		return morse.ToMorse(s), nil
	}
}
