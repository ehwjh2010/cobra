package response

const (
	OK      = 0
	SUCCESS = "Success"

	InvalidParams = 10000
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResult(data interface{}, args ...ResultOpt) *Result {
	result := &Result{
		Code:    OK,
		Message: SUCCESS,
		Data:    data,
	}
	for _, arg := range args {
		arg(result)
	}

	return result
}

type ResultOpt func(response *Result)

func ResultWithCode(code int) ResultOpt {
	return func(response *Result) {
		response.Code = code
	}
}

func ResultWithMessage(msg string) ResultOpt {
	return func(response *Result) {
		response.Message = msg
	}
}

func ResultWithData(data interface{}) ResultOpt {
	return func(response *Result) {
		response.Data = data
	}
}
