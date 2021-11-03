package resource

import (
	"ginLearn/utils"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func LoadMySQL() {
	if Conf.MysqlConfig == nil {
		return
	}

	db, err := utils.InitMySQL(Conf.MysqlConfig)
	if err != nil {
		utils.PanicF("Load mysql failed!, err: %v", err)
	}

	Conn = db
}

func ReleaseMySQL() error {
	return utils.CloseMySQL(Conn)
}
