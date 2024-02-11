package svc

import (
	"gorm.io/gorm"
	"gozerotry/common/init_db"
	"gozerotry/rpc_study/user_grom/models"
	"gozerotry/rpc_study/user_grom/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := init_db.InitGorm(c.Mysql.DataSource)
	err := db.AutoMigrate(&models.UserModel{})
	if err != nil {
		return nil
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
