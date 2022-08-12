// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "go-zero-chat/apps/app/api/internal/handler/user"
	ws "go-zero-chat/apps/app/api/internal/handler/ws"
	"go-zero-chat/apps/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/loginCode",
				Handler: user.LoginCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: user.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: user.SaveHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/token",
				Handler: user.TokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/chat",
				Handler: ws.ChatHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/websocket"),
	)
}