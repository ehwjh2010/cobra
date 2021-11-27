package types

import (
	"fmt"
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

func (r *Result) String() string {
	return fmt.Sprintf("Result(code=%d, message=%s, data=%+v)", r.Code, r.Message, r.Data)
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

func (p *Pageable) String() string {
	return fmt.Sprintf("Pageable(totalCount=%d, totalPage=%d, page=%d, pageSize=%d, rows=%+v, hasNext=%v)",
			p.TotalCount, p.TotalPage, p.Page, p.PageSize, p.Rows, p.HasNext,
		)
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

func NewPageable(rows interface{}, page int, pageSize int, totalCount int64) *Pageable {
	if page <= 0 {
		page = config.DefaultPage
	}

	if pageSize <= 0 {
		pageSize = config.DefaultPageSize
	}

	if totalCount < 0 {
		totalCount = 0
	}

	totalPage := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	hasNext := totalPage > page

	pageable := &Pageable{
		Rows:     rows,
		Page:     page,
		PageSize: pageSize,
		TotalCount: totalCount,
		TotalPage: totalPage,
		HasNext: hasNext,
	}

	return pageable
}
