package class

import (
	"context"

	"github.com/gonhon/go-zero-book/service/user/api/internal/svc"
	"github.com/gonhon/go-zero-book/service/user/api/internal/types"
	"github.com/gonhon/go-zero-book/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClassLogic {
	return &UpdateClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateClassLogic) UpdateClass(req *types.ClassUpdate) (bool, error) {
	// todo: add your logic here and delete this line\
	err := l.svcCtx.ClassModel.Update(l.ctx, &model.Class{
		Code: req.Code,
		Name: req.Name,
		Num:  req.Num,
		Id:   req.Id,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
