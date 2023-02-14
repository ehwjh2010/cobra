package middleware

import (
	"strings"
	"time"

	"github.com/ehwjh2010/viper/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AccessLog 使用ZAP接管GIN相关日志
// timeFormat 时间格式
// utc 是否使用UTC时间
// skipPath 不记录日志的url.
func AccessLog(skipPaths []string, utc bool, timeFormat string) gin.HandlerFunc {
	log.Debugf("Use ginzap middleware")
	skipPaths = append(skipPaths, "/swagger")

	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		echo := true

		for _, skipPath := range skipPaths {
			if strings.HasPrefix(path, skipPath) {
				echo = false
				break
			}
		}

		if echo {
			end := time.Now()
			latency := end.Sub(start).Milliseconds()
			if utc {
				end = end.UTC()
			}

			if len(c.Errors) > 0 {
				// Append error field if this is an erroneous request.
				for _, e := range c.Errors.Errors() {
					log.Errorf(e)
				}
			} else {
				log.Infof(path,
					zap.String("startTime", start.Format(timeFormat)),
					zap.Int("status", c.Writer.Status()),
					zap.String("method", c.Request.Method),
					zap.String("path", path),
					zap.String("query", query),
					zap.String("ip", c.ClientIP()),
					//zap.String("user-agent", c.Request.UserAgent()),
					zap.String("endTime", end.Format(timeFormat)),
					zap.Int64("cost(ms)", latency),
				)
			}
		}
	}
}
