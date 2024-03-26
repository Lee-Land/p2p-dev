package proto

import (
	"encoding/binary"
	"fmt"
	"net"
)

func Parse(conn net.Conn) (*Message, error) {
	sizeBuf := make([]byte, 4)
	n, err := conn.Read(sizeBuf)
	if err != nil {
		return nil, err
	}

	if n < 4 {
		return nil, fmt.Errorf("unexpected %d bytes", n)
	}

	size := binary.BigEndian.Uint32(sizeBuf)

	bodyBuf := make([]byte, size)
	n, err = conn.Read(bodyBuf)
	if err != nil {
		return nil, err
	}

	if uint32(n) < size {
		return nil, fmt.Errorf("unexpected %d bytes", n)
	}

	var msg Message
	msg.Ver = bodyBuf[0]
	msg.Method = Method(binary.BigEndian.Uint16(bodyBuf[1:3]))

	switch msg.Method {
	case Conn:
	case Addr:
	}
	return nil, nil
}
