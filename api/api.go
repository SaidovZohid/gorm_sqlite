package api

import (
	v1 "github.com/SaidovZohid/gorm_sqlite/api/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterOptions struct {
	GormDB *gorm.DB
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		GormDB: opt.GormDB,
	})

	router.POST("/users", handlerV1.CreateUser)
	router.GET("/users/:id", handlerV1.GetUser)
	router.PUT("/users/:id", handlerV1.UpdateUser)
	router.DELETE("/users/:id", handlerV1.DeleteUser)

	return router
}
