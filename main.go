// main.go
package main

import (
	"veilnet/core/node"
	"time"
)

func main() {
	// Генерация ключей (в реальности — AES-256)
	keyA := make([]byte, 32)
	keyB := make([]byte, 32)
	keyC := make([]byte, 32)

	// Узлы
	nodeA := &node.Node{ID: "A", ListenAddr: ":8001", PrivateKey: keyA}
	nodeB := &node.Node{ID: "B", ListenAddr: ":8002", PrivateKey: keyB}
	nodeC := &node.Node{ID: "C", ListenAddr: ":8003", PrivateKey: keyC}

	// Запуск узлов в горутинах
	go nodeA.Start()
	go nodeB.Start()
	go nodeC.Start()

	// Ждём запуска
	time.Sleep(time.Second)

	// Клиент: шифруем "Hello, Veil!" → C → B → A
	payload := []byte("Hello, Veil!")

	// Шифруем "изнутри": сначала для C, потом B, потом A
	layerC, _ := crypto.EncryptOnion(payload, [][]byte{keyC})
	pktC := &packet.Packet{Content: layerC, IsLast: true, SessionID: "123", NextHopID: ":8003"}

	layerB, _ := crypto.EncryptOnion(pktC.Serialize(), [][]byte{keyB})
	pktB := &packet.Packet{Content: layerB, IsLast: false, SessionID: "123", NextHopID: ":8002"}

	layerA, _ := crypto.EncryptOnion(pktB.Serialize(), [][]byte{keyA})
	pktA := &packet.Packet{Content: layerA, IsLast: false, SessionID: "123", NextHopID: ":8001"}

	// Отправляем на первый узел (A)
	conn, _ := network.Dial(":8001")
	conn.Send(pktA.Serialize())
	conn.Close()

	// Держим программу живой
	select {}
}