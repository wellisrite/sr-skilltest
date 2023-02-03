package utilities

import (
	"crypto/rand"
	"encoding/hex"
	"strings"

	"github.com/google/uuid"
)

func RandomHex(number int) (string, error) {
	bytes := make([]byte, number)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func CreateTraceID() string {
	randomHex, _ := RandomHex(5)
	uid := strings.ReplaceAll(uuid.New().String(), "-", "")
	return strings.ToUpper(uid[0:10] + randomHex)
}
