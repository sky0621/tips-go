package model

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

var _ ValueObject[string] = (*UserID)(nil)

type UserID struct {
	value string
}

func (u *UserID) Value() string {
	return u.value
}

func NewUserID(value string) (*UserID, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return nil, errors.New("userID must not be empty")
	}
	return &UserID{value: trimmed}, nil
}

func GenerateNewUserID() (*UserID, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random ID: %w", err)
	}
	// 16バイト -> 32桁のhex string
	return &UserID{value: hex.EncodeToString(bytes)}, nil
}
