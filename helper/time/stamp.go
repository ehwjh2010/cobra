package time

import "time"

// Timestamp 时间戳, 单位为: s.
func Timestamp() int64 {
	return time.Now().Unix()
}

// MsTimestamp 时间戳, 单位为: ms.
func MsTimestamp() int64 {
	return time.Now().UnixMilli()
}

// McTimestamp 时间戳, 单位为: mc.
func McTimestamp() int64 {
	return time.Now().UnixMicro()
}

// NaTimestamp 时间戳, 单位为: na.
func NaTimestamp() int64 {
	return time.Now().UnixNano()
}
