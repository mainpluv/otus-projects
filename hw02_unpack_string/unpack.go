package hw02unpackstring

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			if i == 0 || unicode.IsDigit(rune(s[i-1])) || s[i-1] == '/' {
				return "", ErrInvalidString
			}
			n := int(s[i] - '0')
			var curr, part string
			i--
			if i != 0 && s[i-1] == '/' {
				curr = "/"
			}
			curr += string(s[i])
			for j := 1; j <= n; j++ {
				part += curr
			}
			res = part + res
		} else {
			res = string(s[i]) + res
		}
	}
	return res, nil
}
