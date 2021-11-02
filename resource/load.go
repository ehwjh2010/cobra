package resource

import "ginLearn/utils"

func init() {
	LoadConfig()
	utils.InitLogrus(Conf.Application, Conf.LogConfig)
	utils.InitMySQL(Conf.MysqlConfig)
}
