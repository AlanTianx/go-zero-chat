package user

import (
	"go-zero-chat/apps/app/api/internal/svc"
	"go-zero-chat/apps/app/api/internal/types"
	"go-zero-chat/apps/user/rpc/user"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginCodeLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	limiter *limit.PeriodLimit
}

func NewLoginCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginCodeLogic {
	return &LoginCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginCodeLogic) LoginCode(req *types.LoginCodeReq) (resp *types.LoginCodeResp, err error) {

	v, err := l.svcCtx.UserRpc.LoginCode(l.ctx, &user.LoginCodeReq{
		Phone: req.Phone,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.LoginCodeResp{
		CodeKey: v.CodeKey,
	}, nil
}
