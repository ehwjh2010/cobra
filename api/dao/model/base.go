package model

import "time"

type BaseModel struct {
	ID        int64     `gorm:"column:id;primaryKey;type:bigint(20) unsigned not null auto_increment;comment:主键"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:datetime not null default current_timestamp;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;type:datetime not null default current_timestamp on update current_timestamp;comment:更新时间"`
}
