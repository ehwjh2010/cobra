package enums

type DBType int

// Desc 描述
func (dbType DBType) Desc() string {
	switch dbType {
	case Mysql:
		return "mysql"
	case Postgresql:
		return "postgresql"
	case Sqlite:
		return "sqlite"
	default:
		return ""
	}
}

const (
	Mysql DBType = iota + 1
	Postgresql
	Sqlite
)
