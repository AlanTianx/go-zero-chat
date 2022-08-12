package user

import (
	"go-zero-chat/apps/user/rpc/user"
	"context"
	"encoding/json"
	"github.com/pkg/errors"

	"go-zero-chat/apps/app/api/internal/svc"
	"go-zero-chat/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {

	uId, _ := l.ctx.Value("uid").(json.Number).Int64()
	v, err := l.svcCtx.UserRpc.Detail(l.ctx, &user.DetailReq{Id: uId})
	if err != nil {
		return nil, errors.Wrapf(err, "req%+v", req)
	}

	return &types.UserInfoResp{
		Uuid:     v.Uuid,
		Username: v.Username,
		Phone:    v.Phone,
	}, nil
}
