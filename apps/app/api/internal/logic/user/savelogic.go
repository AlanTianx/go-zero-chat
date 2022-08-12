package user

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"go-zero-chat/apps/user/rpc/user"

	"go-zero-chat/apps/app/api/internal/svc"
	"go-zero-chat/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLogic {
	return &SaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveLogic) Save(req *types.UserSaveReq) (resp *types.UserSaveResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	_, err = l.svcCtx.UserRpc.Save(l.ctx, &user.SaveReq{
		Id:       uid,
		Username: req.Username,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req%+v", req)
	}

	return
}
