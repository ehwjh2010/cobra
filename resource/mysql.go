package resource

import "ginLearn/utils"

var MysqlClient utils.MysqlClient

func LoadMySQL() {
	err := MysqlClient.SetUp(Conf.MysqlConfig)
	if err != nil {
		utils.PanicF("Load mysql failed!, err: %v", err)
	}
}
