// internal/middleware/logger.go

package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"log/slog"
)

func Logger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Обработка запроса
		c.Next()

		// После завершения
		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path

		// Логируем как структурированное сообщение
		logger.Info("HTTP request",
			"method", method,
			"path", path,
			"status", statusCode,
			"latency", latency.Seconds(),
			"client_ip", clientIP,
		)
	}
}
