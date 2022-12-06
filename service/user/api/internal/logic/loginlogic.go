package logic

import (
	"context"
	"strings"
	"time"

	"github.com/gonhon/go-zero-book/common/base"
	"github.com/gonhon/go-zero-book/service/user/api/internal/svc"
	"github.com/gonhon/go-zero-book/service/user/api/internal/types"
	"github.com/gonhon/go-zero-book/service/user/model"

	"github.com/form3tech-oss/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

/*
*
生成token
*/
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["ext"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))

}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, base.NewDefaultError("参数错误")
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, req.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, base.NewDefaultError("用户名不存在")
	default:
		return nil, err
	}
	if req.Password != userInfo.Password {
		return nil, base.NewDefaultError("密码错误")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret,
		now,
		l.svcCtx.Config.Auth.AccessExpire,
		userInfo.Id)
	if err != nil {
		return nil, err
	}
	resp = &types.LoginReply{
		Id:           userInfo.Id,
		Name:         userInfo.Name,
		Gender:       userInfo.Gender,
		AccessToken:  token,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}
	return
}
