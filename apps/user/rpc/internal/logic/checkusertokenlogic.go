package logic

import (
	"context"
	"github.com/pkg/errors"
	"time"

	"go-zero-chat/apps/user/rpc/internal/svc"
	"go-zero-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserTokenLogic {
	return &CheckUserTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserTokenLogic) CheckUserToken(in *user.CheckTokenReq) (*user.CheckTokenRes, error) {
	v, err := l.svcCtx.UserTokensModel.FindOne(l.ctx, in.GetTokenMd5())
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", in)
	}

	if v.ExpireTime < time.Now().Unix() {
		return nil, errors.Wrapf(errors.New("token已经失效"), "req: %+v", in)
	}

	u, err := l.svcCtx.UserModel.FindOne(l.ctx, v.UserId)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", in)
	}

	return &user.CheckTokenRes{
		Id:       v.Id,
		Username: u.Name,
		Uuid:     u.Uuid,
	}, nil
}
