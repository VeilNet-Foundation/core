// integrations/ipfs.go
func MountIPFSOverVeil() {
	// Все запросы к IPFS идут через .veil-сеть
	ipfs.DHT = "/dns/ipfs.veil"
}