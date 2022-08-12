package user

import (
	"go-zero-chat/pkg/result"
	"net/http"

	"go-zero-chat/apps/app/api/internal/logic/user"
	"go-zero-chat/apps/app/api/internal/svc"
	"go-zero-chat/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewSaveLogic(r.Context(), svcCtx)
		resp, err := l.Save(&req)
		result.HttpResult(r, w, resp, err)
	}
}
