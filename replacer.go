package main

import "strings"

// TokenReplacer replaces any tokens in a 'text' with
// the values in 'replacements'.
type TokenReplacer interface {
	replace(replacements map[string]string, text string) string
}

type defaultReplacer struct {
}

func (r defaultReplacer) replace(replacements map[string]string, text string) string {
	t := text

	for k, v := range replacements {
		t = strings.Replace(t, "["+k+"]", v, -1)
	}

	return t
}
