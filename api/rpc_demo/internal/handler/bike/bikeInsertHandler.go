package bike

import (
	"net/http"

	"go-zero-demo/api/rpc_demo/internal/logic/bike"
	"go-zero-demo/api/rpc_demo/internal/svc"
	"go-zero-demo/api/rpc_demo/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func BikeInsertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BikeInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := bike.NewBikeInsertLogic(r.Context(), svcCtx)
		resp, err := l.BikeInsert(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
