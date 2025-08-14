// crypto/utils.go
package crypto

import "crypto/sha256"

func GenerateID(pubKey []byte) string {
	hash := sha256.Sum256(pubKey)
	return fmt.Sprintf("%x", hash[:8]) // короткий ID
}