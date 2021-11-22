package response

import (
	"github.com/ehwjh2010/cobra/config"
	"math"
)

const (
	OK      = 0
	SUCCESS = "Success"

	InvalidParams = 10000
)

type Result struct {
	//Code 业务状态码
	Code int `json:"code" example:"0" swaggertype:"integer"`

	//Message 信息
	Message string `json:"message" example:"Success" swaggertype:"string"`

	//Data 数据
	Data interface{} `json:"data"`
}

type Pageable struct {
	//TotalCount 总数量
	TotalCount int64 `json:"totalCount"`

	//TotalPage 总页数
	TotalPage int `json:"totalPage"`

	//Page 当前页数
	Page int `json:"page"`

	//PageSize 每页数量
	PageSize int `json:"pageSize"`

	//Rows 记录
	Rows interface{} `json:"rows"`

	//HasNext 是否还有下一页
	HasNext bool `json:"hasNext"`
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

type PageableOpt func(pageable *Pageable)

//PageableWithPage 设置页数, 页数必须大于0
func PageableWithPage(page int) PageableOpt {
	return func(pageable *Pageable) {
		pageable.Page = page
	}
}

//PageableWithPageSize 设置每页数量, 页数必须大于0
func PageableWithPageSize(pageSize int) PageableOpt {
	return func(pageable *Pageable) {
		pageable.PageSize = pageSize
	}
}

//PageableWithTotalCount 设置总数量
func PageableWithTotalCount(totalCount int64) PageableOpt {
	return func(pageable *Pageable) {
		pageable.TotalCount = totalCount
	}
}

//NewPageable 默认页数为1, 每页数量15
func NewPageable(rows interface{}, args ...PageableOpt) *Pageable {
	pageable := &Pageable{
		Rows:     rows,
		Page:     config.DefaultPage,
		PageSize: config.DefaultPageSize,
	}

	for _, arg := range args {
		arg(pageable)
	}

	if pageable.PageSize > 0 {
		totalPage := int(math.Ceil(float64(pageable.TotalCount) / float64(pageable.PageSize)))
		pageable.TotalPage = totalPage
	}

	pageable.HasNext = pageable.TotalPage > pageable.Page

	return pageable
}
