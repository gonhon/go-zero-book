package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gonhon/go-zero-book/service/search/api/internal/svc"
	"github.com/gonhon/go-zero-book/service/search/api/internal/types"
	"github.com/gonhon/go-zero-book/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchReply, err error) {
	// todo: add your logic here and delete this line
	logx.Info("search...")
	//获取token中的userId
	userId := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId:%s", userId)
	id, err := userId.Int64()
	if err != nil {
		return nil, err
	}

	res, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	if res != nil {
		logx.Info("res:", res)
	}

	resp = &types.SearchReply{
		Name:  req.Name,
		Count: 100,
	}
	return
}
