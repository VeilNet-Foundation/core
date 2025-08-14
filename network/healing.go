// network/healing.go
func (n *Node) RunHealthCheck() {
  for _, peer := range n.Peers {
    if !ping(peer) {
      n.reportFailure(peer)
      n.rebuildTunnels()
    }
  }
}
func (n *Node) HealNetwork() {
	peers := n.GetPeers()
	
	for _, p := range peers {
		if !n.Ping(p) {
			n.reportFailure(p)
			n.rebuildTunnelsExcluding(p)
			go n.BroadcastOutage(p)
		}
	}
	
	// Каждые 30 сек
	time.Sleep(30 * time.Second)
	n.HealNetwork()
}