package verror

import "strings"

type MultiErr struct {
	Errs []error
}

func (m *MultiErr) Error() string {
	if m.IsEmpty() {
		return ""
	}

	var tmp []string
	for _, err := range m.Errs {
		tmp = append(tmp, err.Error())
	}

	result := strings.Join(tmp, "\n")
	return result
}

// AddErr 添加错误
func (m *MultiErr) AddErr(args ...error) {
	for _, err := range args {
		if err == nil {
			continue
		}

		m.Errs = append(m.Errs, err)
	}
}

// IsEmpty 是否为空
func (m *MultiErr) IsEmpty() bool {
	if m == nil || len(m.Errs) == 0 {
		return true
	}

	return false
}

// IsNotEmpty 是否为空
func (m *MultiErr) IsNotEmpty() bool {
	return !m.IsNotEmpty()
}

// AsStdErr 转换为标准库错误
func (m *MultiErr) AsStdErr() error {
	if m.IsEmpty() {
		return nil
	}

	return m
}
