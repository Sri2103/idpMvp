// pkg/utils/utils.go
package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateID(prefix string) string {
	bytes := make([]byte, 4)
	_, _ = rand.Read(bytes)
	return prefix + "-" + hex.EncodeToString(bytes)
}
