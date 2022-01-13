package global

var NullBytes = []byte("null")

const NullStr = "null"

const SwaggerAPIUrl = "/swagger/index.html"

const HomeShortCut = "~" //类Unix系统home路径的短符号

const (
	DefaultPage     = 1  //默认页数
	DefaultPageSize = 15 //默认每页数据
)

const (
	Chinese = "cn" //中文
	English = "en" //英文
)

const (
	ContentType        = "Content-Type"
	ContentDisposition = "Content-Disposition"
	UserAgent          = "ViperRequests/1.1.0"

	JsonType      = "application/json; charset=utf-8"
	FormType      = "application/x-www-form-urlencoded; charset=utf-8"
	MultiFormType = "multipart/form-data; charset=utf-8; boundary=--801fe6c28526e72589981c923d518232--"
)
