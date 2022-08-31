package ws

import (
	"context"
	"github.com/kataras/neffos"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/hash"
	"go-zero-chat/apps/user/rpc/user"

	"go-zero-chat/apps/app/api/internal/svc"
	"go-zero-chat/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	OnChatEvent = "chat"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (l *ChatLogic) WithCtx(ctx context.Context) *ChatLogic {
	c := *l
	c.ctx = ctx
	return &c
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

func (l *ChatLogic) OnChat(nsConn *neffos.NSConn, msg neffos.Message) error {
	// todo 解析 msg 进行具体业务处理 ---

	return nil
}
