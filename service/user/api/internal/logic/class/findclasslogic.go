package class

import (
	"context"

	"github.com/gonhon/go-zero-book/service/user/api/internal/svc"
	"github.com/gonhon/go-zero-book/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindClassLogic {
	return &FindClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindClassLogic) FindClass(req *types.FindClassReq) (resp []types.ClassReply, err error) {
	// todo: add your logic here and delete this line

	return
}
