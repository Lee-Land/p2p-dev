package proto

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

func Unpack(buffer []byte) (*Message, error) {
	if len(buffer) < headerLength {
		return nil, fmt.Errorf("buffer is not enough %d bytes", headerLength)
	}
	headerBuffer := buffer[:headerLength]
	return &Message{
		Header: Header{
			Ver:    binary.LittleEndian.Uint16(headerBuffer[:2]),
			Method: Method(binary.LittleEndian.Uint16(headerBuffer[2:])),
		},
		Payload: buffer[headerLength:],
	}, nil
}

func Pack(msg *Message) ([]byte, error) {
	ver := make([]byte, 2)
	method := make([]byte, 2)
	binary.LittleEndian.PutUint16(ver, msg.Ver)
	binary.LittleEndian.PutUint16(method, uint16(msg.Method))

	header := append(ver, method...)
	header = append(header, msg.Payload...)
	return header, nil
}

func Read(conn net.Conn) (*Message, error) {
	buffer := make([]byte, 4)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	if n < 4 {
		return nil, errors.New("buffer is not enough 4 bytes")
	}

	msgLen := binary.LittleEndian.Uint32(buffer)

	buffer = make([]byte, msgLen)
	n, err = conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	if uint32(n) < msgLen {
		return nil, fmt.Errorf("buffer is not enough %d bytes", msgLen)
	}

	return Unpack(buffer)
}
