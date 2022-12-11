package class

import (
	"context"

	"github.com/gonhon/go-zero-book/service/user/api/internal/svc"
	"github.com/gonhon/go-zero-book/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClassLogic {
	return &DeleteClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteClassLogic) DeleteClass(req *types.ClassReq) (resp bool, err error) {
	// todo: add your logic here and delete this line

	return
}
