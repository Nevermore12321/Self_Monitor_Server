package svc

import (
	"api/internal/config"
	model "github.com/Nevermore12321/Self_Monitor_Server/service/api-gateway/model/user_info"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserInfoModel

}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserInfoModel(conn, c.CacheRedis),
	}
}
