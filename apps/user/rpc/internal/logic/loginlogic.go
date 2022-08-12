package logic

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go-zero-chat/pkg/helper"

	"github.com/zeromicro/go-zero/core/hash"

	"go-zero-chat/apps/user/rpc/internal/svc"
	"go-zero-chat/apps/user/rpc/model"
	"go-zero-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginRes, error) {
	// 验证验证码
	v, err := l.svcCtx.BizRedis.Get(in.CodeKey)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", in)
	}
	if !bytes.Equal([]byte(v), hash.Md5([]byte(fmt.Sprintf("%s%s", in.Phone, in.Code)))) {
		return nil, errors.Wrapf(errors.New("验证码错误"), "req:%+v", in)
	}
	if in.Phone == "" {
		return nil, errors.Wrapf(errors.New("参数错误-phone不能为空"), "req:%+v", in)
	}
	defer l.svcCtx.BizRedis.Del(in.CodeKey)

	uuidx := uuid.NewV4()

	uv, err := l.svcCtx.UserModel.FirstOrCreateByPhone(l.ctx, model.Users{
		Uuid:  uuidx.String(),
		Name:  fmt.Sprintf("V-%s", helper.RandStr(5)),
		Phone: in.Phone,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req:%+v", in)
	}

	// 查找user
	return &user.LoginRes{
		Id:       uv.Id,
		Username: uv.Name,
		Phone:    uv.Phone,
		Uuid:     uv.Uuid,
	}, nil
}
