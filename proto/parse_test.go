package proto

import (
	"fmt"
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
