package model

import (
	"ginLearn/types"
)

//BaseModel 表通用字段
type BaseModel struct {
	ID        int64          `gorm:"column:id;primaryKey;type:bigint(20) unsigned not null auto_increment;comment:主键" json:"id"`
	CreatedAt types.NullTime `gorm:"column:created_at;autoCreateTime;type:datetime not null default current_timestamp;comment:创建时间" json:"createdAt"`
	UpdatedAt types.NullTime `gorm:"column:updated_at;autoUpdateTime;type:datetime not null default current_timestamp on update current_timestamp;comment:更新时间" json:"updatedAt"`
}
