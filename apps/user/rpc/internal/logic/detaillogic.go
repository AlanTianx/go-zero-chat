package logic

import (
	"context"
	"github.com/pkg/errors"

	"go-zero-chat/apps/user/rpc/internal/svc"
	"go-zero-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *user.DetailReq) (*user.LoginRes, error) {

	u, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", in)
	}
	// 查找user
	return &user.LoginRes{
		Id:       u.Id,
		Username: u.Name,
		Phone:    u.Phone,
		Uuid:     u.Uuid,
	}, nil
}
