package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/models"
	"user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	// 添加我们需要放入的数据模型
	UserModel models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 创建合适的数据库连接
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		// 利用数据库连接将真实的数据模型注入
		UserModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}
