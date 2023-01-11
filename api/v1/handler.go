package v1

import (
	"gorm.io/gorm"
)

type handlerV1 struct {
	gormDB *gorm.DB
}

type HandlerV1Options struct {
	GormDB *gorm.DB
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		gormDB: options.GormDB,
	}
}
