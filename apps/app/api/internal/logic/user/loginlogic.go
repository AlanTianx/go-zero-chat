package user

import (
	"go-zero-chat/apps/app/api/internal/svc"
	"go-zero-chat/apps/app/api/internal/types"
	"go-zero-chat/apps/user/rpc/user"
	"go-zero-chat/pkg/jwtx"
	"context"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.LoginResp, err error) {
	v, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Phone:   req.Phone,
		CodeKey: req.CodeKey,
		Code:    req.Code,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req:%+v", req)
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, v.Id)

	return &types.LoginResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil
}
