package resource

import (
	"ginLearn/utils"
	"gorm.io/gorm"
	"log"
)

var Conn *gorm.DB

func LoadMySQL() {
	if Conf.MysqlConfig == nil {
		return
	}

	err := utils.InitMySQL(Conf.MysqlConfig, Conn)
	if err != nil {
		log.Fatalf("Init mysql failed!, err: %v", err)
	}
}

func ReleaseMySQL() error {
	return utils.CloseMySQL(Conn)
}
