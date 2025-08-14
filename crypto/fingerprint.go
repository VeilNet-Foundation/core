// crypto/fingerprint.go
func ObfuscateAs(protocol string, data []byte) []byte {
  switch protocol {
  case "quic":
    return quicEmulate(data)
  case "bittorrent":
    return bittorrentNoise(data)
  default:
    return addRandomPadding(data)
  }
}