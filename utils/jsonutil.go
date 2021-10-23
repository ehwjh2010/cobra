package utils

import jsoniter "github.com/json-iterator/go"

func JsonMarshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}
