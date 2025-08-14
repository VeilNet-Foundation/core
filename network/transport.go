// network/transport.go
package network

import "net"

type Conn interface {
	Send([]byte) error
	Receive() ([]byte, error)
	Close() error
}

type TCPConn struct {
	conn net.Conn
}

func (t *TCPConn) Send(data []byte) error {
	_, err := t.conn.Write(data)
	return err
}

func (t *TCPConn) Receive() ([]byte, error) {
	var sizeBuf = make([]byte, 4)
	_, err := t.conn.Read(sizeBuf)
	if err != nil {
		return nil, err
	}
	size := int(binary.BigEndian.Uint32(sizeBuf))
	buf := make([]byte, size)
	_, err = t.conn.Read(buf)
	return buf, err
}

func (t *TCPConn) Close() error {
	return t.conn.Close()
}

func Listen(addr string) (net.Listener, error) {
	return net.Listen("tcp", addr)
}

func Dial(addr string) (*TCPConn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &TCPConn{conn: conn}, nil
}