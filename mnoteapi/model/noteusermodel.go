package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	noteUserFieldNames          = builder.RawFieldNames(&NoteUser{})
	noteUserRows                = strings.Join(noteUserFieldNames, ",")
	noteUserRowsExpectAutoSet   = strings.Join(stringx.Remove(noteUserFieldNames, "`create_time`", "`update_time`"), ",")
	noteUserRowsWithPlaceHolder = strings.Join(stringx.Remove(noteUserFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheNoteUserIdPrefix = "cache:noteUser:id:"
)

type (
	NoteUserModel interface {
		Insert(data *NoteUser) (sql.Result, error)
		FindOneByNamePassword(name, password string) (*NoteUser, error)
		FindOne(id int64) (*NoteUser, error)
		Update(data *NoteUser) error
		Delete(id int64) error
	}

	defaultNoteUserModel struct {
		sqlc.CachedConn
		table string
	}

	NoteUser struct {
		Id       int64  `db:"id"`
		Name     string `db:"name"`
		Password string `db:"password"`
		Nickname string `db:"nickname"`
		Identity string `db:"identity"`
	}
)

func NewNoteUserModel(conn sqlx.SqlConn, c cache.CacheConf) NoteUserModel {
	return &defaultNoteUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`note_user`",
	}
}

func (m *defaultNoteUserModel) Insert(data *NoteUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, noteUserRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.Name, data.Password, data.Nickname, data.Identity)

	return ret, err
}

func (m *defaultNoteUserModel) FindOneByNamePassword(name, password string) (*NoteUser, error) {
	// TODO encrypt password
	enPass := password

	noteUserIdKey := fmt.Sprintf("%s%v%v", cacheNoteUserIdPrefix, name, enPass)
	var resp NoteUser
	err := m.QueryRow(&resp, noteUserIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `name` = ? and `password` = ? limit 1", noteUserRows, m.table)
		return conn.QueryRow(v, query, name, enPass)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNoteUserModel) FindOne(id int64) (*NoteUser, error) {
	noteUserIdKey := fmt.Sprintf("%s%v", cacheNoteUserIdPrefix, id)
	var resp NoteUser
	err := m.QueryRow(&resp, noteUserIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", noteUserRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNoteUserModel) Update(data *NoteUser) error {
	noteUserIdKey := fmt.Sprintf("%s%v", cacheNoteUserIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, noteUserRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Password, data.Nickname, data.Identity, data.Id)
	}, noteUserIdKey)
	return err
}

func (m *defaultNoteUserModel) Delete(id int64) error {

	noteUserIdKey := fmt.Sprintf("%s%v", cacheNoteUserIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, noteUserIdKey)
	return err
}

func (m *defaultNoteUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheNoteUserIdPrefix, primary)
}

func (m *defaultNoteUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", noteUserRows, m.table)
	return conn.QueryRow(v, query, primary)
}
