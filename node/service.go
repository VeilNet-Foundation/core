// node/service.go
func (n *Node) HostService(serviceID string, handler func([]byte)) {
	dht.Put(serviceID, dht.Entry{
		NodeAddr:  n.ListenAddr,
		PublicKey: n.PublicKey,
		TTL:       time.Now().Add(1 * time.Hour).Unix(),
	})
	// Создаём входящий туннель
	inTunnel := &tunnel.Tunnel{
		ID:      "in-" + serviceID,
		Hops:    []string{n.ID, "middle", "guard"},
		Reverse: true,
	}
	tunnelManager.Add(inTunnel)
}