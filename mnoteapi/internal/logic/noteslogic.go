package logic

import (
	"context"
	"fmt"

	"github.com/moyrne/mnote/mnoteapi/internal/svc"
	"github.com/moyrne/mnote/mnoteapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	userInfo, err := l.svcCtx.UserFilter(l.ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println(userInfo)
	return
}
