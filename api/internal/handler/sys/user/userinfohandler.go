package user

import (
	"net/http"

	"go-zero-demo/api/common/response"
	"go-zero-demo/api/internal/logic/sys/user"
	"go-zero-demo/api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, resp, err)
	}
}
