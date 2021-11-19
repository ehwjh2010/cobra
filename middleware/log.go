package middleware

import (
	"fmt"
	"github.com/ehwjh2010/cobra/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

//GinZap 使用ZAP接管GIN相关日志
//timeFormat 时间格式
//utc 是否使用UTC时间
//skipPath 不记录日志的url
func GinZap(conf *MiddleConfig) gin.HandlerFunc {
	fmt.Println("Use ginzap middleware")
	skipPaths := make(map[string]bool, len(conf.SkipPaths))
	for _, path := range conf.SkipPaths {
		skipPaths[path] = true
	}

	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		if _, ok := skipPaths[path]; !ok {
			end := time.Now()
			latency := end.Sub(start)
			if conf.UTC {
				end = end.UTC()
			}

			if len(c.Errors) > 0 {
				// Append error field if this is an erroneous request.
				for _, e := range c.Errors.Errors() {
					log.Error(e)
				}
			} else {
				log.Infol(path,
					zap.Int("status", c.Writer.Status()),
					zap.String("method", c.Request.Method),
					zap.String("path", path),
					zap.String("query", query),
					zap.String("ip", c.ClientIP()),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.String("time", end.Format(conf.TimeFormat)),
					zap.Duration("latency", latency),
				)
			}
		}
	}
}
