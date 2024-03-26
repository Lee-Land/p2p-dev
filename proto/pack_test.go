package proto

import (
	"fmt"
	"p2p-dev/assert"
	"testing"
	"unsafe"
)

type Example struct {
	A uint8
	B uint16
	C uint32
}

func TestMemory(t *testing.T) {
	example := Example{
		A: 42,
		B: 1,
		C: 2,
	}

	// 打印结构体的大小和字段在内存中的偏移量
	fmt.Printf("Size of Example struct: %d bytes\n", unsafe.Sizeof(example))
	fmt.Printf("Offset of A: %d bytes\n", unsafe.Offsetof(example.A))
	fmt.Printf("Offset of B: %d bytes\n", unsafe.Offsetof(example.B))
	fmt.Printf("Offset of C: %d bytes\n", unsafe.Offsetof(example.C))
}

func TestUnpack(t *testing.T) {
	buffer, err := Pack(&Message{
		Header: Header{
			Ver:    1,
			Method: Conn,
		},
		Payload: []byte{0x01, 0x02, 0x03, 0x04},
	})
	assert.NoError(t, err)

	msg, err := Unpack(buffer)
	assert.NoError(t, err)

	assert.Equals(t, msg.Ver, 1)
	assert.Equals(t, msg.Method, Conn)
}
