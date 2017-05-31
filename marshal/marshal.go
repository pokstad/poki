package marshal

import "github.com/pokstad/poki"

type MarshalPost interface {
	Marshal(poki.Post) ([]byte, error)
}

type UnmarshalPost interface {
	Unmarshal(data []byte, dst *poki.Post) error
}
