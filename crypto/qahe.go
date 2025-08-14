// crypto/qahe.go
func GenerateKeyFromEntropy() [32]byte {
  raw := collectEntropyFromHardware() // microphone, camera noise
  strength := ai.EvaluateEntropy(raw) // локальная модель
  if strength < 0.9 {
    panic("Слишком предсказуемая энтропия")
  }
  return sha256.Sum256(raw)
}