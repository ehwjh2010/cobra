package resource

import "ginLearn/utils"

var DBClient *utils.DBClient

func LoadDB() {

	client, err := utils.InitDB(Conf.DBConfig)

	if err != nil {
		Logger.PanicF("Load mysql failed!, err: %v", err)
	}

	DBClient = client
}
