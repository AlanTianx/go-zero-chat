package user

import (
	"context"
	"encoding/json"
	"go-zero-chat/apps/user/rpc/user"

	"go-zero-chat/apps/app/api/internal/svc"
	"go-zero-chat/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TokenLogic {
	return &TokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenLogic) Token(req *types.UserTokenReq) (resp *types.UserTokenResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	v, err := l.svcCtx.UserRpc.GetUserToken(l.ctx, &user.GetUserTokenReq{Id: uid})
	if err != nil {
		return nil, err
	}

	return &types.UserTokenResp{
		Token:      v.GetToken(),
		ExpireTime: v.GetExpireTime(),
	}, nil
}
