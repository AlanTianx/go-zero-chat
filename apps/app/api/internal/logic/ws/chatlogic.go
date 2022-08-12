package ws

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/hash"
	"go-zero-chat/apps/user/rpc/user"

	"go-zero-chat/apps/app/api/internal/svc"
	"go-zero-chat/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type UserInfo struct {
	Uuid     string
	Username string
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) CheckToken(req *types.WebsocketReq) (resp *UserInfo, err error) {
	v, err := l.svcCtx.UserRpc.CheckUserToken(l.ctx, &user.CheckTokenReq{TokenMd5: hash.Md5Hex([]byte(req.Token))})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &UserInfo{
		Uuid:     v.Uuid,
		Username: v.Username,
	}, nil
}
