package ws

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
	stackexchange "github.com/kataras/neffos/stackexchange/redis"
	"go-zero-chat/apps/app/api/internal/logic/ws"
	"go.opentelemetry.io/otel/trace"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-chat/apps/app/api/internal/svc"
	"go-zero-chat/apps/app/api/internal/types"
)

const (
	Namespace = "chat"
)

var server = neffos.New(gorilla.Upgrader(websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// todo 允许所有的CORS 跨域请求，正式环境可以关闭
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}), neffos.Namespaces{
	Namespace: neffos.Events{
		neffos.OnNamespaceConnected: func(nsConn *neffos.NSConn, msg neffos.Message) error {
			// 服务器回复客户端
			nsConn.Emit(msg.Event, []byte((fmt.Sprintf("欢迎%s加入聊天室", nsConn.Conn.ID()))))

			// 广播给其他客户端 我进入了聊天室
			nsConn.Conn.Server().Broadcast(nsConn, neffos.Message{
				Namespace: msg.Namespace,
				//Room:      msg.Room,
				Event: msg.Event,
				Body:  []byte((fmt.Sprintf("欢迎%s加入聊天室", nsConn.Conn.ID()))),
				To:    "",
			})

			return nil
		},
		neffos.OnNamespaceDisconnect: func(nsConn *neffos.NSConn, msg neffos.Message) error {
			// 广播给其他客户端 我离开了聊天室
			nsConn.Conn.Server().Broadcast(nsConn, neffos.Message{
				Namespace: msg.Namespace,
				//Room:      msg.Room,
				Event: msg.Event,
				Body:  []byte((fmt.Sprintf("%s离开了聊天室", nsConn.Conn.ID()))),
				To:    "",
			})
			return nil
		},
		// todo 引入 Room 时可以使用
		//neffos.OnRoomJoin: func(nsConn *neffos.NSConn, message neffos.Message) error {
		//	return nil
		//},
		//neffos.OnRoomLeft: func(nsConn *neffos.NSConn, message neffos.Message) error {
		//	return nil
		//},
		ws.OnChatEvent: func(nsConn *neffos.NSConn, msg neffos.Message) error {
			myCtx := GetContext(nsConn.Conn)
			// todo ctx可以作为后续具体业务func的链路追踪
			ctx := trace.ContextWithSpanContext(context.TODO(), myCtx.spanCtx)

			return myCtx.l.WithCtx(ctx).OnChat(nsConn, msg)

			// ps gorilla 消息发送说明
			// 广播消息给某个客户端
			//nsConn.Conn.Server().Broadcast(nsConn, neffos.Message{
			//	Namespace: msg.Namespace, // todo 你也可以广播给其他namespace
			//	Room:      msg.Room,      // todo 你也可以广播给其他Room
			//	Event:     msg.Event,     // todo 你也可以广播给其他Event
			//	Body:      []byte((fmt.Sprintf("i am %s", nsConn.Conn.ID()))),
			//	To:        "A", // 假设A是某个客户端的nsConn.Conn.ID()
			//})

			// 广播消息给其他客户端-自己不接收此消息
			//nsConn.Conn.Server().Broadcast(nsConn, neffos.Message{
			//	Namespace: msg.Namespace,
			//	Room:      msg.Room,
			//	Event:     msg.Event,
			//	Body:      []byte((fmt.Sprintf("i am %s", nsConn.Conn.ID()))),
			//	To:        "",
			//})

			// 广播消息给所有客户端
			//nsConn.Conn.Server().Broadcast(nil, neffos.Message{
			//	Namespace: msg.Namespace,
			//	Room:      msg.Room,
			//	Event:     msg.Event,
			//	Body:      []byte((fmt.Sprintf("i am %s", nsConn.Conn.ID()))),
			//	To:        "",
			//})
		},
	},
})

// WsServerInit 初始化ws-server设置 main 函数中调用
func WsServerInit(svcCtx *svc.ServiceContext) {
	// 设置消息广播为同步 异步可能丢失消息
	server.SyncBroadcaster = true

	// 设置网络交互  以支持集群 该操作会覆盖 server.SyncBroadcaster设置 让它无效。
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
	svcCtx  *svc.ServiceContext
	spanCtx trace.SpanContext // spanCtx 可以在后续ws交互中进行链路追踪
	w       http.ResponseWriter
	r       *http.Request
	l       *ws.ChatLogic
}

// GetContext returns the Iris Context from a websocket connection.
func GetContext(c *neffos.Conn) *MyWsContext {
	if sw, ok := c.Socket().(*socketWrapper); ok {
		return sw.ctx
	}
	return nil
}

// ChatHandler api接口 处理升级ws
func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WebsocketReq

		// 解析api参数
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// 鉴权---
		l := ws.NewChatLogic(r.Context(), svcCtx)
		u, err := l.CheckToken(&req)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		r.Form.Add("username", u.Username)

		// 给ws的连接生成一个id
		idGen := func(w http.ResponseWriter, r *http.Request) string {
			return r.FormValue("username")
		}

		// todo 提取出 spanCtx 后面websocket交互中每次生成新的context 加入spanCtx
		spanCtx := trace.SpanContextFromContext(r.Context())

		// 升级成ws
		_, err = server.Upgrade(w, r, func(socket neffos.Socket) neffos.Socket {
			return &socketWrapper{
				Socket: socket,
				ctx: &MyWsContext{
					svcCtx:  svcCtx,
					spanCtx: spanCtx,
					w:       w,
					r:       r,
					l:       l,
				},
			}
		}, idGen)

		if err != nil {
			httpx.Error(w, err)
		}
	}
}

// ServerBroadcast 供其他线程试用 服务器发消息给client 直接使用chatServer广播
func ServerBroadcast(exceptSender fmt.Stringer, msg ...neffos.Message) {
	server.Broadcast(exceptSender, msg...)
}

// GetClientConn Waring 非线程安全、勿频繁使用！！！
func GetClientConn(id string) *neffos.Conn {
	all := server.GetConnections()
	return all[id]
}
