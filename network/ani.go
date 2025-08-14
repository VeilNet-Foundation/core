// network/ani.go
func StartDecoyFlows() {
  for {
    go func() {
      path := BuildRandomPath()
      payload := generateRandomQuery()
      SendMultipathEncrypted(payload, path, true) // isDecoy = true
    }()
    time.Sleep(randTime(100*time.Millisecond, 5*time.Second))
  }
}