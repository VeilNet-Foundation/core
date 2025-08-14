// message/deferred.go
func SendLater(msg []byte, recipient string) {
  delay := time.Hour * time.Duration(1 + rand.Intn(23))
  schedule(func() {
    SendMultipath(msg, recipient)
  }, delay)
}