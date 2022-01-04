package handler

import (
	"net/http"

	"github.com/moyrne/mnote/mnoteapi/internal/logic"
	"github.com/moyrne/mnote/mnoteapi/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func NotesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewNotesLogic(r.Context(), svcCtx)
		resp, err := l.Notes()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
