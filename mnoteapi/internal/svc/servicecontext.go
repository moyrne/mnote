package svc

import (
	"context"

	"github.com/moyrne/mnote/mnoteapi/internal/config"
	"github.com/moyrne/mnote/mnoteapi/model"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	NoteUserModel model.NoteUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		NoteUserModel: model.NewNoteUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}

type UserInfo struct {
	UserID int64 `json:"user_id"`
}

func (s ServiceContext) UserFilter(ctx context.Context) (*UserInfo, error) {
	userID, ok := ctx.Value("userID").(int64)
	if !ok {
		return nil, errors.WithStack(ErrUserNotLogin)
	}

	return &UserInfo{UserID: userID}, nil
}
