package class

import (
	"context"

	"github.com/gonhon/go-zero-book/common/base"
	"github.com/gonhon/go-zero-book/service/user/api/internal/svc"
	"github.com/gonhon/go-zero-book/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdLogic {
	return &GetByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByIdLogic) GetById(req *types.ClassPath) (resp *types.ClassReply, err error) {
	// todo: add your logic here and delete this line
	class, e := l.svcCtx.ClassModel.FindOne(l.ctx, req.Id)
	if e != nil {
		logx.Error(e)
		return nil, base.NewCodeError(500, "查询错误")
	}
	resp = &types.ClassReply{
		Id:   class.Id,
		Name: class.Name,
		Code: class.Code,
		Num:  class.Num,
	}
	return
}
