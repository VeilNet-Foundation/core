// crypto/onion.go
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// EncryptOnion создаёт "лук" из пакета, шифруя слоями
func EncryptOnion(payload []byte, keys [][]byte) ([]byte, error) {
	encrypted := payload
	for i := len(keys) - 1; i >= 0; i-- {
		block, err := aes.NewCipher(keys[i])
		if err != nil {
			return nil, err
		}
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return nil, err
		}
		nonce := make([]byte, gcm.NonceSize())
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return nil, err
		}
		encrypted = gcm.Seal(nonce, nonce, encrypted, nil)
	}
	return encrypted, nil
}

// DecryptOneLayer расшифровывает один слой
func DecryptOneLayer(layer []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(layer) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := layer[:nonceSize], layer[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}