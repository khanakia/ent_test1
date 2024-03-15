package server_handler

import (
	"lace/nlog"
	"saas/gen/ent"

	"lace/ginserver"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	// GormDB gormdb.GormDB
	Client *ent.Client
	Logger nlog.Logger
	Server ginserver.Server
}

type Handler struct {
	Config
}

func (pkg Handler) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func authValidateHandler(c *gin.Context) {
	c.JSON(200, gin.H{"code": 200})
}

func New(config Config) Handler {
	pkg := Handler{
		Config: config,
	}

	// db := pkg.Config.GormDB.DB

	router := config.Server.Router

	authorized := router.Group("/p")
	// authorized.Use(middleware.JwtMiddleware(pkg.Config.Client))
	authorized.GET("/auth/validate", authValidateHandler)

	return pkg
}
