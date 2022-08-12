package user

import (
	"go-zero-chat/pkg/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-chat/apps/app/api/internal/logic/user"
	"go-zero-chat/apps/app/api/internal/svc"
	"go-zero-chat/apps/app/api/internal/types"
)

func TokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewTokenLogic(r.Context(), svcCtx)
		resp, err := l.Token(&req)
		result.HttpResult(r, w, resp, err)
	}
}
