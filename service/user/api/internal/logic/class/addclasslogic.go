package class

import (
	"context"

	"github.com/gonhon/go-zero-book/common/base"
	"github.com/gonhon/go-zero-book/service/user/api/internal/svc"
	"github.com/gonhon/go-zero-book/service/user/api/internal/types"
	"github.com/gonhon/go-zero-book/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassLogic {
	return &AddClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddClassLogic) AddClass(req *types.ClassReq) (bool, error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.ClassModel.Insert(l.ctx, &model.Class{
		Num:  req.Num,
		Name: req.Name,
		Code: req.Code,
	})
	if err != nil {
		logx.Error(err)
		return false, base.NewCodeError(500, "数据插入错误")
	}
	logx.Infof("res:%v", res)
	return true, nil
}
