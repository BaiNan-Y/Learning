package logic

import (
	"context"

	"imooc.com/traning/user/rpc/internal/svc"
	"imooc.com/traning/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 定义rpc方法
func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	if u, ok := users[in.Id]; ok {
		return &user.GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}, nil
	}
	return &user.GetUserResp{}, nil
}
