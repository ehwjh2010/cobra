package resource

func Release() []error {
	mysqlErr := ReleaseMySQL()

	//...其他需要处理的数据

	logFileErr := ReleaseLogrus()

	errs := []error{mysqlErr, logFileErr}
	return errs
}
