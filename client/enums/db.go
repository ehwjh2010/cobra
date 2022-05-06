package enums

//go:generate stringer -type DBType -linecomment
type DBType int

const (
	Mysql      DBType = iota + 1 // mysql
	Postgresql                   // postgresql
	Sqlite                       // sqlite
)

//go:generate stringer -type DBCbWhen
type DBCbWhen int

const (
	Before DBCbWhen = iota + 1
	After
)

//go:generate stringer -type OperateType
type OperateType int

const (
	Register OperateType = iota + 1
	Replace
	Remove
)
