package cserde

import (
	"github.com/vmihailenco/msgpack/v5"
)

// Better Than JSON Serialization & De-serialization Library.

// Serialize a struct into MessagePack Binary Data
func Se(v interface{}) ([]byte, error) {
	b, err := msgpack.Marshal(&v)

	if err != nil {
		panic(err)
	}

	return b, nil
}

func Deserialize() {}
