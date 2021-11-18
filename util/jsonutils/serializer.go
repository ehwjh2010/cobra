package jsonutils

import (
	"github.com/json-iterator/go"
)

//Marshal 对应json.Marshal
func Marshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}

//Unmarshal 对应json.Unmarshal
func Unmarshal(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}
