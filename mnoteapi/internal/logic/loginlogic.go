package logic

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/moyrne/mnote/mnoteapi/internal/svc"
	"github.com/moyrne/mnote/mnoteapi/internal/types"
	"github.com/pkg/errors"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginRequest) (resp *types.LoginResponse, err error) {
	user, err := l.svcCtx.NoteUserModel.CheckPassword(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	token, err := l.getJWTToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, user.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Status:   "ok",
		Message:  "success",
		ID:       user.Id,
		Name:     user.Name,
		Nickname: user.Nickname,
		Token:    token,
		Expire:   now + accessExpire,
	}, nil
}

func (l *LoginLogic) getJWTToken(secretKey string, iat, seconds, userID int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userID"] = userID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	t, err := token.SignedString(secretKey)
	return t, errors.WithStack(err)
}
