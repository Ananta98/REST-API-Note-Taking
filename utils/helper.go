package utils

import (
	"net/mail"
	"net/url"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func IsValidUrl(str string) bool {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}
	return true
}

func IsValidEmail(str string) bool {
	_, err := mail.ParseAddress(str)
	if err != nil {
		return false
	}
	return true
}
