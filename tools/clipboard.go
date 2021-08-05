package tools

import "github.com/atotto/clipboard"

func ReadClip() (string, error) {
	return clipboard.ReadAll()
}

func WriteClip(content string) error {
	return clipboard.WriteAll(content)
}
