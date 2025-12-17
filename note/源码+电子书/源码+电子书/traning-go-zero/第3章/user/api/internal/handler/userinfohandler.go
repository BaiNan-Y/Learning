package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"imooc.com/traning/user/api/internal/logic"
	"imooc.com/traning/user/api/internal/svc"
	"imooc.com/traning/user/api/internal/types"
)

func userinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
