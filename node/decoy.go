// node/decoy.go
func (n *Node) ServeDecoy(w http.ResponseWriter, r *http.Request) {
  // Имитация торрент-трекера
  w.Header().Set("Content-Type", "text/plain")
  w.Write([]byte("d5:filesd..."))
}
func (n *Node) StartDecoyServer() {
	http.HandleFunc("/announce", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("d14:failure reason23:Invalid announce request"))
	})
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(generateRandomHTML())
	})
	
	go http.ListenAndServe(":80", nil)
}