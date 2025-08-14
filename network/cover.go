// network/cover.go
func SendWithCover(pkt *Packet) {
  // Дополняем до 1024 байт
  padded := padTo(pkt.Serialize(), 1024)
  
  // Отправляем с задержкой
  time.Sleep(randomJitter(100*time.Millisecond, 2*time.Second))
  
  // С вероятностью 30% — фиктивный пакет
  if rand.Float32() < 0.3 {
    padded = generateRandomPacket(1024)
  }
  
  transport.Send(padded)
}