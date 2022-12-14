// Code generated by goctl. DO NOT EDIT.
package types

type LoginCodeReq struct {
	Phone string `form:"phone"`
}

type LoginCodeResp struct {
	CodeKey string `json:"codeKey"`
}

type UserLoginReq struct {
	Phone   string `form:"phone"`
	Code    string `form:"code"`
	CodeKey string `form:"code_key"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

type UserSaveReq struct {
	Username string `form:"username"`
}

type UserSaveResp struct {
}

type UserTokenReq struct {
}

type UserTokenResp struct {
	Token      string `json:"token"`
	ExpireTime int64  `json:"expireTime"`
}

type WebsocketReq struct {
	Token string `form:"token"`
}

type WebsocketResp struct {
}
