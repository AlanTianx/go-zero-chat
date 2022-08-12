package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/hash"
	"go-zero-chat/apps/user/rpc/internal/svc"
	"go-zero-chat/apps/user/rpc/model"
	"go-zero-chat/apps/user/rpc/user"
	"go-zero-chat/pkg/helper"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTokenLogic {
	return &GetUserTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserTokenLogic) GetUserToken(in *user.GetUserTokenReq) (*user.GetUserTokenRes, error) {
	token := helper.RandStr(128)
	tokenMd5 := hash.Md5Hex([]byte(token))
	et := time.Now().Add(time.Minute * 5).Unix()
	err := l.svcCtx.UserTokensModel.Insert(l.ctx, &model.UserTokens{
		UserId:     in.GetId(),
		Token:      token,
		TokenKey:   tokenMd5,
		Status:     1,
		ExpireTime: et,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req:%+v", in)
	}

	return &user.GetUserTokenRes{
		Token:      token,
		ExpireTime: et,
	}, nil
}
