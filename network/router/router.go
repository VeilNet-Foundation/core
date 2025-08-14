// router/router.go
package router

import (
	"veilnet/core/crypto"
	"veilnet/core/packet"
)

type Route struct {
	Nodes  []string
	Keys   [][]byte
}

func BuildMultipathRoutes(dest string, allNodes []Node, paths int) []Route {
	var routes []Route
	for i := 0; i < paths; i++ {
		hops := randomPath(allNodes, 3) // guard → middle → exit
		keys := extractKeys(hops)
		routes = append(routes, Route{Nodes: hops, Keys: keys})
	}
	return routes
}

func SendMultipath(payload []byte, routes []Route) {
	for _, route := range routes {
		encrypted, _ := crypto.EncryptOnion(payload, route.Keys)
		pkt := &packet.Packet{
			Content:   encrypted,
			NextHopID: route.Nodes[0],
			SessionID: genSession(),
			IsLast:    len(route.Nodes) == 1,
		}
		forward(pkt, route.Nodes[0])
	}
}