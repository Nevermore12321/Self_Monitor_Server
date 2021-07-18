package svc

import (
	"api/internal/config"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	model "github.com/Nevermore12321/Self_Monitor_Server/service/api-gateway/model/user_info"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserInfoModel

}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel: userModel.NewUserInfoModel(conn, c.CacheRedis),
	}
}
