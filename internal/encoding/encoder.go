package encoding

import (
	"errors"
	"sync"
)

type Encoder interface {
	Encode(v interface{}) ([]byte, error)
}

type EncoderRegistry struct {
	encoders map[string]Encoder

	mu sync.Mutex
}

func (e *EncoderRegistry) RegisterEncoder(format string, enc Encoder) error {
	if _, ok := e.encoders[format]; ok {
		return errors.New("format already registered")
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.encoders[format] = enc

	return nil
}

func (e *EncoderRegistry) Encode(format string, v interface{}) ([]byte, error) {
	encoder, ok := e.encoders[format]
	if !ok {
		return nil, errors.New("encoder not found for format")
	}

	return encoder.Encode(v)
}
