package logic

import (
	"go-zero-chat/apps/user/rpc/internal/svc"
	"go-zero-chat/apps/user/rpc/model"
	"go-zero-chat/apps/user/rpc/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLogic {
	return &SaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveLogic) Save(in *user.SaveReq) (*user.SaveRes, error) {
	err := l.svcCtx.UserModel.Update(l.ctx, &model.Users{
		Id:   in.Id,
		Name: in.Username,
	})
	if err != nil {
		// todo 回写消息队列、重新尝试修改
	}

	return &user.SaveRes{}, nil
}
