package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
)

func ComputeHMAC256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func ComputeHMAC512(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func IsMatchedHMAC256(signature string, message string, secret string) bool {
	return ComputeHMAC256(message, secret) == signature
}

func IsMatchedHMAC512(signature string, message string, secret string) bool {
	return ComputeHMAC512(message, secret) == signature
}
