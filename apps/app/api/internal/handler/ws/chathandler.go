package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
	stackexchange "github.com/kataras/neffos/stackexchange/redis"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-chat/apps/app/api/internal/logic/ws"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-chat/apps/app/api/internal/svc"
	"go-zero-chat/apps/app/api/internal/types"
)

var server = neffos.New(gorilla.Upgrader(websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// todo 允许所有的CORS 跨域请求，正式环境可以关闭
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}), neffos.Namespaces{
	"chat": neffos.Events{
		neffos.OnNamespaceConnected: func(nsConn *neffos.NSConn, msg neffos.Message) error {
			logx.Info("websocket Connected success:", nsConn.Conn.ID())
			return nil
		},
		neffos.OnNamespaceDisconnect: func(nsConn *neffos.NSConn, msg neffos.Message) error {
			logx.Info("websocket disConnected")
			return nil
		},
		"chat": func(nsConn *neffos.NSConn, message neffos.Message) error {
			ctx := GetContext(nsConn.Conn)
			logx.Info(ctx.svcCtx.Config)
			nsConn.Conn.Server().Broadcast(nsConn, neffos.Message{
				Namespace: "AgentPreStart",
				Room:      "",
				Event:     "chat",
				Body:      []byte((fmt.Sprintf("i am %s", nsConn.Conn.ID()))),
				To:        "",
			})
			return nil
		},
	},
})

func WsServerInit(svcCtx *svc.ServiceContext) {
	// 设置消息广播为同步 异步可能丢失消息
	server.SyncBroadcaster = true

	// 设置网络交互  以支持集群 该操作会覆盖 server.SyncBroadcaster 设置让它失效。
	cfg := stackexchange.Config{
		Addr:     svcCtx.Config.BizRedis.Host,
		Password: svcCtx.Config.BizRedis.Pass,
	}
	redisStg, err := stackexchange.NewStackExchange(cfg, "chat")
	if err != nil {
		panic("err: without redisStackExchange")
	}
	err1 := server.UseStackExchange(redisStg)
	if err1 != nil {
		panic("err: without redisStackExchange")
	}
}

type socketWrapper struct {
	neffos.Socket
	ctx *MyWsContext
}

type MyWsContext struct {
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
	l      *ws.ChatLogic
}

// GetContext returns the Iris Context from a websocket connection.
func GetContext(c *neffos.Conn) *MyWsContext {
	if sw, ok := c.Socket().(*socketWrapper); ok {
		return sw.ctx
	}
	return nil
}

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WebsocketReq

		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := ws.NewChatLogic(r.Context(), svcCtx)
		u, err := l.CheckToken(&req)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		r.Form.Add("username", u.Username)

		idGen := func(w http.ResponseWriter, r *http.Request) string {
			return r.FormValue("username")
		}

		_, err = server.Upgrade(w, r, func(socket neffos.Socket) neffos.Socket {
			return &socketWrapper{
				Socket: socket,
				ctx: &MyWsContext{
					svcCtx: svcCtx,
					w:      w,
					r:      r,
					l:      l,
				},
			}
		}, idGen)

		if err != nil {
			httpx.Error(w, err)
		}
	}
}
