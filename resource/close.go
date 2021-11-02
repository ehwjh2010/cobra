package resource

import "ginLearn/utils"

func Close() []error {
	mysqlErr := utils.CloseMySQL()

	//...其他需要处理的数据

	logFileErr := utils.CloseLogFile()

	errs := []error{mysqlErr, logFileErr}
	return errs
}
