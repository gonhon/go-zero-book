package class

import (
	"context"

	"github.com/gonhon/go-zero-book/common/base"
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

func (l *DeleteClassLogic) DeleteClass(req *types.ClassPath) (bool, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.ClassModel.Delete(l.ctx, req.Id)
	if err != nil {
		return false, base.NewCodeError(500, "删除失败")
	}
	return true, err
}
