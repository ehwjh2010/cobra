package verror

import (
	"errors"
)

var (
	CastStrErr      = errors.New("convert str failed")
	CastStrSliceErr = errors.New("convert str slice failed")
	CastBoolErr     = errors.New("convert bool failed")
	CastIntErr      = errors.New("convert int failed")
	CastInt32Err    = errors.New("convert int32 failed")
	CastInt64Err    = errors.New("convert int64 failed")
	CastFloat64Err  = errors.New("convert float64 failed")
)
