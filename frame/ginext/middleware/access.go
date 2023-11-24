package middleware

import (
	"strings"
	"time"

	"github.com/ehwjh2010/viper/log"
	"github.com/gin-gonic/gin"
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
				log.Infof("%s, startTime: %s, status: %d, method: %s, query: %s, IP: %s, endTime: %s, cost(ms): %d ms",
					path, start.Format(timeFormat), c.Writer.Status(), c.Request.Method, query, c.ClientIP(), end.Format(timeFormat), latency)
			}
		}
	}
}
