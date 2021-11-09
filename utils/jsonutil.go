package utils

import jsoniter "github.com/json-iterator/go"

//JsonMarshal 对应json.Marshal
func JsonMarshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}

//JsonUnmarshal 对应json.Unmarshal
func JsonUnmarshal(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}
