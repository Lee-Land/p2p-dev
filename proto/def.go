package proto

type Method uint16

const (
	_ Method = iota
	Conn
	Addr
)

type Message struct {
	Ver     uint8
	Method  Method
	Payload []byte
}

type Address struct {
	IPv4 uint32
	Port uint16
}

type MsgConn struct {
	SrcAddr Address
	DstAddr Address
}

type MsgAddr struct {
	Addresses []*Address
}

type MsgUser struct {
	Payload []byte
}
