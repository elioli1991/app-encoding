package encoding

import (
	"strings"

	"github.com/elioli1991/app-infra/abstract"
)

var registerCodecs = make(map[string]abstract.Codec)

// RegisterCodec register codec
func RegisterCodec(codec abstract.Codec) {
	if codec == nil || codec.Name() == "" {
		return
	}
	registerCodecs[strings.ToLower(codec.Name())] = codec
}

// GetCodec get codec by content type name
func GetCodec(name string) abstract.Codec {
	return registerCodecs[strings.ToLower(name)]
}
