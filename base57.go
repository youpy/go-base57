package base57

import (
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v3"
)

type base57 struct {
	encoder shortuuid.Encoder
}

func New() *base57 {
	return &base57{shortuuid.DefaultEncoder}
}

func (b base57) Encode(u uuid.UUID) string {
	return reverseString(b.encoder.Encode(u))
}

func (b base57) Decode(u string) (uuid.UUID, error) {
	return b.encoder.Decode(reverseString(u))
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
