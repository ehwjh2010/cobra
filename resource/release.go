package resource

var errs []error

func Release() []error {
	mysqlErr := ReleaseMySQL()
	addErr(mysqlErr)

	//...其他需要处理的数据

	logFileErr := ReleaseLogrus()
	addErr(logFileErr)

	return errs
}

func addErr(err error) {
	if err != nil {
		errs = append(errs, err)
	}
}
