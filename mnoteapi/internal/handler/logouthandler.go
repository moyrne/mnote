package handler

import (
	"net/http"

	"github.com/moyrne/mnote/mnoteapi/internal/logic"
	"github.com/moyrne/mnote/mnoteapi/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
