package main

import (
	"crypto/rand"
	"fmt"
)

func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
