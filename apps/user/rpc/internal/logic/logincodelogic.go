package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/hash"
	"go-zero-chat/apps/sms/sms"
	"go-zero-chat/pkg/helper"
	"strconv"

	"go-zero-chat/apps/user/rpc/internal/svc"
	"go-zero-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var noSendPhone = map[string]int{"13838182466": 1, "13838182465": 2}

func NewLoginCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginCodeLogic {
	return &LoginCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginCodeLogic) LoginCode(in *user.LoginCodeReq) (*user.LoginCodeRes, error) {
	code := helper.RandNum(4)
	codeKey := helper.RandStr(15)
	//if bytes.Equal([]byte(in.Phone), []byte("13838182466")) {
	if _, ok := noSendPhone[in.Phone]; ok {
		code = 1234
	}
	_, err := l.svcCtx.BizRedis.SetnxEx(codeKey, hash.Md5Hex([]byte(fmt.Sprintf("%s%v", in.Phone, code))), 320)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", in)
	}

	// 发送短信
	_, err = l.svcCtx.SmsRpc.SendSms(l.ctx, &sms.SendSmsReq{
		Action: "login",
		Phone:  in.Phone,
		Msg:    strconv.Itoa(code),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", in)
	}

	return &user.LoginCodeRes{
		Code:    strconv.Itoa(code),
		CodeKey: codeKey,
	}, nil
}
