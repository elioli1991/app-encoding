package proto

import (
	"errors"

	encoding "github/elioli1991/app-encoding"

	"google.golang.org/protobuf/proto"
)

const (
	Name = "proto"
)

var ErrorProtoMessageType = errors.New("error proto message type")

func init() {
	encoding.RegisterCodec(codec{})
}

type codec struct{}

func (codec) Marshal(v any) ([]byte, error) {
	if _, ok := v.(proto.Message); !ok {
		return nil, ErrorProtoMessageType
	}
	return proto.Marshal(v.(proto.Message))
}

func (codec) Unmarshal(data []byte, v any) error {
	if _, ok := v.(proto.Message); !ok {
		return ErrorProtoMessageType
	}
	return proto.Unmarshal(data, v.(proto.Message))
}

func (codec) Name() string {
	return Name
}
