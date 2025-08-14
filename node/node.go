// node/node.go
package node

import (
	"veilnet/core/crypto"
	"veilnet/core/network"
	"veilnet/core/packet"
)

type Node struct {
	ID       string
	ListenAddr string
	PrivateKey []byte
}

func (n *Node) Start() error {
	ln, err := network.Listen(n.ListenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go n.handleConnection(conn)
	}
}

func (n *Node) handleConnection(conn net.Conn) {
	tcpConn := &network.TCPConn{conn: conn}
	data, err := tcpConn.Receive()
	if err != nil {
		return
	}

	pkt, err := packet.Deserialize(data)
	if err != nil {
		return
	}

	// Расшифровать один слой
	decrypted, err := crypto.DecryptOneLayer(pkt.Content, n.PrivateKey)
	if err != nil {
		// Это не наш слой — просто пересылаем
		forwardToNext(pkt.NextHopID, data)
		return
	}

	// Если мы последний — обрабатываем как сервис
	if pkt.IsLast {
		go n.handleFinal(decrypted)
	} else {
		// Иначе распаковываем пакет и пересылаем дальше
		nextPkt, err := packet.Deserialize(decrypted)
		if err != nil {
			return
		}
		forwardToNext(nextPkt.NextHopID, nextPkt.Serialize())
	}
}

func (n *Node) handleFinal(payload []byte) {
	println("Final node received:", string(payload))
}

func forwardToNext(addr string, data []byte) {
	conn, err := network.Dial(addr)
	if err != nil {
		return
	}
	defer conn.Close()
	conn.Send(data)
}