package util

import b64 "encoding/base64"

// Base64EncodeToString ...
func Base64EncodeToString(text string) string {
	encoded := b64.StdEncoding.EncodeToString([]byte(text))
	return encoded
}
