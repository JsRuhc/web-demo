package routes

import (
	"net/http"
	"web_demo/logger"
	"web_demo/settings"

	"github.com/gin-gonic/gin"
)

func SetUp(cfg *settings.AppConfig) *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, cfg.Level)
	})
	return r
}
