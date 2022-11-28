package user

import (
	"net/http"

	"go-zero-demo/api/rpc_demo/internal/logic/user"
	"go-zero-demo/api/rpc_demo/internal/svc"
	"go-zero-demo/api/rpc_demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RpcUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRes
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewRpcUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.RpcUserInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
