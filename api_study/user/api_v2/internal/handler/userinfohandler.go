package handler

import (
	"gozerotry/common/response"
	"net/http"

	"gozerotry/api_study/user/api_v2/internal/logic"
	"gozerotry/api_study/user/api_v2/internal/svc"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(r, w, resp, err)
	}
}
