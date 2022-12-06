package logic

import (
	"context"

	"github.com/gonhon/go-zero-book/service/user/rpc/internal/svc"
	"github.com/gonhon/go-zero-book/service/user/rpc/types/user"

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

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReply, error) {
	res, err := l.svcCtx.UserMudel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &user.UserInfoReply{
		Id:     res.Id,
		Name:   res.Name,
		Number: res.Number,
		Gender: res.Gender,
	}, nil
}
