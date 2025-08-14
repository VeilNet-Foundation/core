// crypto/qrog.go
func HybridEncrypt(data []byte, pqPubKey, classicalPubKey []byte) []byte {
  // 1. Генерируем сеансовый ключ
  sharedSecret := kyber.Encapsulate(pqPubKey)
  
  // 2. Добавляем классический ключ (на случай уязвимости PQ)
  hybridKey := hash(sharedSecret + ecdh.Shared(classicalPubKey))
  
  // 3. Шифруем данные
  return aes256gcm.Seal(data, hybridKey)
}
// Post-Quantum + Classical Hybrid
func QROG_Encrypt(data []byte, pqPub, classicalPub []byte) []byte {
	// 1. Квантово-устойчивый обмен
	pqShared := kyber.Encapsulate(pqPub)
	
	// 2. Классический (резерв)
	classicalShared := ecdh.ComputeSecret(classicalPub, privateKey)
	
	// 3. Гибридный ключ
	hybrid := sha3.Sum256(append(pqShared, classicalShared...))
	
	// 4. Шифрование
	return chacha20poly1305.Seal(data, hybrid[:32])
}