package yaml

import "gopkg.in/yaml.v2"

type Codec struct{}

func (Codec) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (Codec) Decode(b []byte, v interface{}) error {
	return yaml.Unmarshal(b, v)
}
