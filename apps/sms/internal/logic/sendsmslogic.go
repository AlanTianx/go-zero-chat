package logic

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"go-zero-chat/apps/sms/internal/svc"
	"go-zero-chat/apps/sms/sms"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var noSendPhone = map[string]int{"13838182466": 1, "13838182465": 2}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *sms.SendSmsReq) (*sms.SendSmsResp, error) {
	// todo: add your logic here and delete this line

	resp := &sms.SendSmsResp{
		Code: "ok",
		Msg:  "发送成功",
	}

	if _, ok := noSendPhone[in.Phone]; ok {
		return resp, nil
	}

	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", l.svcCtx.Config.Sms.KeyId, l.svcCtx.Config.Sms.KeySecret)
	if err != nil {
		resp.Code = "failed"
		resp.Msg = err.Error()
		return resp, err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = in.Phone
	request.SignName = "聊天室"

	switch in.Action {
	case "login":
		l.captcha(in, request)
	}

	r, err := client.SendSms(request)

	if err != nil {
		resp.Code = "failed"
		resp.Msg = err.Error()
		return resp, err
	}

	if r != nil && !bytes.Equal([]byte("OK"), []byte(r.Code)) {
		resp.Code = "failed"
		resp.Msg = r.Message
		return resp, errors.New(r.Message)
	}

	return resp, nil
}

// 短信验证码目标
func (l *SendSmsLogic) captcha(in *sms.SendSmsReq, request *dysmsapi.SendSmsRequest) {
	request.TemplateCode = "*****"
	request.TemplateParam = fmt.Sprintf("{\"code\":%s}", in.Msg)
}
