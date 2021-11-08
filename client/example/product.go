package example

import (
	"ginLearn/utils"
	"gopkg.in/guregu/null.v4"
)

type Product struct {
	ID         int64        `json:"id"`
	CreatedAt  utils.BJTime `json:"createdAt"`
	UpdatedAt  utils.BJTime `json:"updatedAt"`
	Name       string       `json:"name"`
	Price      float64      `json:"price"`
	TotalCount int64        `json:"totalCount"`
	Brand      null.String  `json:"brand"`
}

func NewProduct() *Product {
	return &Product{}
}
