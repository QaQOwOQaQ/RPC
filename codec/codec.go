package codec

import "io"

type Header struct {
	ServiceMethod string
	Seq           uint64
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implement
)

var NewCodecFunMap map[Type]NewCodecFunc

func init() {
	NewCodecFunMap = make(map[Type]NewCodecFunc)
	NewCodecFunMap[GobType] = NewGobCodec
}
