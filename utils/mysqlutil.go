package utils

import (
	"ginLearn/client/setting"
	"gorm.io/gorm"
)

type MysqlClient struct {
	gormDB *gorm.DB
}

func (c *MysqlClient) SetUp(mysqlConfig *setting.MysqlConfig) error {
	gormDB, err := InitMySQLWithGorm(mysqlConfig)
	if err != nil {
		return err
	}

	c.gormDB = gormDB
	return nil
}

func (c *MysqlClient) Close() error {
	return CloseMySQLWithGorm(c.gormDB)
}
