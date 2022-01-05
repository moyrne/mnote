package svc

import (
	"github.com/moyrne/mnote/mnoteapi/internal/config"
	"github.com/moyrne/mnote/mnoteapi/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
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
