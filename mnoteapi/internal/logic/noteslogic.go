package logic

import (
	"context"

	"github.com/moyrne/mnote/mnoteapi/internal/svc"
	"github.com/moyrne/mnote/mnoteapi/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NotesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotesLogic(ctx context.Context, svcCtx *svc.ServiceContext) NotesLogic {
	return NotesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotesLogic) Notes() (resp *types.StatusResponse, err error) {
	// todo: add your logic here and delete this line
	l.ctx.Value("userID")
	return
}
