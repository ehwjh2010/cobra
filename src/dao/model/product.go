package model

import (
	"gopkg.in/guregu/null.v4"
)

type Product struct {
	BaseModel  `gorm:"embedded"`
	Name       string      `gorm:"column:name;type:varchar(128);not null;default:'';index:idx_name_total_count;comment:商品名称" json:"name"`
	Price      float64     `gorm:"column:price;type:decimal(10,4);not null;default:0;comment:商品价格" json:"price"`
	TotalCount int64       `gorm:"column:total_count;type:int unsigned;not null;default:0;index:idx_name_total_count;comment:商品库存" json:"-"`
	Brand      null.String `gorm:"column:brand;type:varchar(128);comment:品牌" json:"brand"`
}

func NewProduct() *Product {
	return &Product{}
}

//TableName 指定Product结构体对应的数据表为product
func (p Product) TableName() string {
	return "product"
}
