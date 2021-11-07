package model

import (
	"database/sql"
)

type Product struct {
	BaseModel  `gorm:"embedded"`
	Name       string         `gorm:"column:name;type:varchar(128);not null;default:'';index:idx_name_total_count;comment:商品名称"`
	Price      float64        `gorm:"column:price;type:decimal(10,4);not null;default:0;comment:商品价格"`
	TotalCount int64          `gorm:"column:total_count;type:int unsigned;not null;default:0;index:idx_name_total_count;comment:商品库存"`
	Brand      sql.NullString `gorm:"column:brand;type:varchar(128);comment:品牌"`
}

func NewProduct() *Product {
	return &Product{}
}

//TableName 指定Product结构体对应的数据表为product
func (p Product) TableName() string {
	return "product"
}
