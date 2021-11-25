package middleware

import (
	"github.com/ehwjh2010/cobra/util/object"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CorsConfig struct {
	AllowOrigins []string //允许哪些源 源: 协议+域名+端口 -> http://example.com

	AllowMethods []string //允许哪些请求方式

	AllowHeaders []string //允许哪些请求头

	AllowCredentials bool //是否允许传输Cookie

	ExposeHeaders []string //请求方可以拿到哪些请求头

	MaxAge time.Duration //本次预检请求的有效期

	AllowWildcard bool
}

type CorsOpt func(config *CorsConfig)

func OriginOpt(origin ...string) CorsOpt {
	return func(config *CorsConfig) {
		config.AllowOrigins = append(config.AllowOrigins, origin...)
	}
}

func MethodOpt(method ...string) CorsOpt {
	return func(config *CorsConfig) {
		config.AllowMethods = append(config.AllowMethods, method...)
	}
}

func HeaderOpt(header ...string) CorsOpt {
	return func(config *CorsConfig) {
		config.AllowHeaders = append(config.AllowHeaders, header...)
	}
}

func CookieOpt(allow bool) CorsOpt {
	return func(config *CorsConfig) {
		config.AllowCredentials = allow
	}
}

func ExHeaderOpt(header ...string) CorsOpt {
	return func(config *CorsConfig) {
		config.ExposeHeaders = append(config.ExposeHeaders, header...)
	}
}

func MaxAgeOpt(maxAge time.Duration) CorsOpt {
	return func(config *CorsConfig) {
		config.MaxAge = maxAge
	}
}

func Cors(args ...CorsOpt) gin.HandlerFunc {

	config := &CorsConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodHead,
			http.MethodHead,
			http.MethodOptions,
			http.MethodDelete,
		},
		AllowHeaders: []string{
			"Authorization",
			"Content-Length",
			"X-CSRF-Token",
			"Token",
			"Session",
			"Accept",
			"Origin",
			"Host",
			"Connection",
			"Accept-Encoding",
			"Accept-Language",
			"DNT",
			"Keep-Alive",
			"User-Agent",
			"X-Requested-With",
			"If-Modified-Since",
			"Cache-Control",
			"Content-Type",
		},
		AllowCredentials: true,
		ExposeHeaders: []string{
			"Content-Length",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Cache-Control",
			"Content-Language",
			"Content-Type",
			"Expires",
			"Last-Modified",
		},
		MaxAge:        time.Hour * 24,
		AllowWildcard: true,
	}

	for _, arg := range args {
		arg(config)
	}

	c := &cors.Config{}

	object.CopyProperties(config, c)

	return cors.New(*c)
}
