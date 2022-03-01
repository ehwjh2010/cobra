package cast

import "go.uber.org/zap/buffer"

func Str2Bytes(v string) []byte {
	var buff buffer.Buffer

	_, _ = buff.WriteString(v)

	return buff.Bytes()
}
