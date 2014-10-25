package utils

import (
	"crypto/rand"
	"fmt"

	"github.com/russross/blackfriday"
)

func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func ConvertMarkdownToHtml(markdown string) string {
	return string(blackfriday.MarkdownBasic([]byte(markdown)))
}
