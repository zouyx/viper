package encoding

import (
	"errors"
	"sync"
)

type Decoder interface {
	Decode(b []byte, v interface{}) error
}

type DecoderRegistry struct {
	encoders map[string]Decoder

	mu sync.Mutex
}

func (e *DecoderRegistry) RegisterDecoder(format string, enc Decoder) error {
	if _, ok := e.encoders[format]; ok {
		return errors.New("format already registered")
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.encoders[format] = enc

	return nil
}

func (e *DecoderRegistry) Decode(format string, b []byte, v interface{}) error {
	encoder, ok := e.encoders[format]
	if !ok {
		return errors.New("encoder not found for format")
	}

	return encoder.Decode(b, v)
}
