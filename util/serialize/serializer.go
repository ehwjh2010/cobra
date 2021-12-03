package serialize

import (
	"github.com/json-iterator/go"
)

//Marshal 对应json.Marshal
func Marshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}

//MarshalStr 对应json.Marshal
func MarshalStr(v interface{}) (string, error) {
	b, err := jsoniter.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

//Unmarshal 对应json.Unmarshal
func Unmarshal(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}

//UnmarshalStr 对应json.Unmarshal
func UnmarshalStr(data string, v interface{}) error {
	tmp := []byte(data)
	return jsoniter.Unmarshal(tmp, v)
}
