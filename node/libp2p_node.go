// node/libp2p_node.go
package node

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

type LibP2PNode struct {
	Host host.Host
}

func NewLibP2PNode(listenPort int) (*LibP2PNode, error) {
	h, err := libp2p.New(
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", listenPort)),
	)
	if err != nil {
		return nil, err
	}

	node := &LibP2PNode{Host: h}
	node.setupStreamHandler()

	return node, nil
}

func (n *LibP2PNode) setupStreamHandler() {
	n.Host.SetStreamHandler("/veilnet/packet/1.0.0", func(stream network.Stream) {
		// Читаем пакет
		data, err := io.ReadAll(stream)
		if err != nil {
			stream.Reset()
			return
		}
		// Обрабатываем
		go n.handlePacket(data)
		stream.Close()
	})
}

func (n *LibP2PNode) ConnectToPeer(addr string) error {
	peerInfo, err := peer.AddrInfoFromString(addr)
	if err != nil {
		return err
	}
	return n.Host.Connect(context.Background(), *peerInfo)
}