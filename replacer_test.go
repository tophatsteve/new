package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleReplacement(t *testing.T) {
	input := "This is [name] text"
	replacements := map[string]string{
		"name": "test",
	}
	replacer := defaultReplacer{}

	assert.Equal(t, "This is test text", replacer.replace(replacements, input), "[name] should be replaced by test")
}

func TestMultipleReplacementSameKey(t *testing.T) {
	input := "[name] should say '[name]'"
	replacements := map[string]string{
		"name": "test",
	}
	replacer := defaultReplacer{}

	assert.Equal(t, "test should say 'test'", replacer.replace(replacements, input), "[name] should be replaced by test")
}

func TestMultipleReplacementMultipleKeys(t *testing.T) {
	input := "This is [name] text, with [value] value."
	replacements := map[string]string{
		"name":  "test",
		"value": "replaced",
	}
	replacer := defaultReplacer{}

	assert.Equal(t, "This is test text, with replaced value.", replacer.replace(replacements, input), "tokens should be replaced")
}
