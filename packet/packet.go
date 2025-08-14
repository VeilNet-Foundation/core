// packet/packet.go
package packet

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Packet struct {
	Content   []byte
	NextHopID string
	SessionID string
	IsLast    bool
}

// Serialize упаковывает пакет
func (p *Packet) Serialize() []byte {
	var data []byte
	data = append(data, p.packString(p.NextHopID)...)
	data = append(data, p.packString(p.SessionID)...)
	data = append(data, boolToByte(p.IsLast))
	data = append(data, p.packBytes(p.Content)...)
	return data
}

// Deserialize распаковывает пакет
func Deserialize(data []byte) (*Packet, error) {
	r := &reader{ data, offset: 0}
	nextHop, err := r.readString()
	if err != nil {
		return nil, err
	}
	session, err := r.readString()
	if err != nil {
		return nil, err
	}
	isLast, err := r.readBool()
	if err != nil {
		return nil, err
	}
	content, err := r.readBytes()
	if err != nil {
		return nil, err
	}
	return &Packet{
		NextHopID: nextHop,
		SessionID: session,
		IsLast:    isLast,
		Content:   content,
	}, nil
}

func (p *Packet) packString(s string) []byte {
	return append(p.packUint32(uint32(len(s))), []byte(s)...)
}

func (p *Packet) packBytes(b []byte) []byte {
	return append(p.packUint32(uint32(len(b))), b...)
}

func (p *Packet) packUint32(n uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, n)
	return b
}

func boolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

type reader struct {
	data   []byte
	offset int
}

func (r *reader) readString() (string, error) {
	n, err := r.readUint32()
	if err != nil {
		return "", err
	}
	if r.offset+int(n) > len(r.data) {
		return "", io.ErrUnexpectedEOF
	}
	s := string(r.data[r.offset : r.offset+int(n)])
	r.offset += int(n)
	return s, nil
}

func (r *reader) readBytes() ([]byte, error) {
	n, err := r.readUint32()
	if err != nil {
		return nil, err
	}
	if r.offset+int(n) > len(r.data) {
		return nil, io.ErrUnexpectedEOF
	}
	b := make([]byte, n)
	copy(b, r.data[r.offset:r.offset+int(n)])
	r.offset += int(n)
	return b, nil
}

func (r *reader) readUint32() (uint32, error) {
	if r.offset+4 > len(r.data) {
		return 0, io.ErrUnexpectedEOF
	}
	n := binary.BigEndian.Uint32(r.data[r.offset:])
	r.offset += 4
	return n, nil
}

func (r *reader) readBool() (bool, error) {
	if r.offset+1 > len(r.data) {
		return false, io.ErrUnexpectedEOF
	}
	b := r.data[r.offset]
	r.offset++
	return b == 1, nil
}