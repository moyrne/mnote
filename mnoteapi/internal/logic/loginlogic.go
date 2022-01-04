package logic

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/moyrne/mnote/mnoteapi/internal/svc"
	"github.com/moyrne/mnote/mnoteapi/internal/types"

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

func (l *LoginLogic) Login(req types.LoginRequest) (resp *types.StatusResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

func (l *LoginLogic) getJWTToken(secretKey string, iat, seconds, userID int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["user_id"] = userID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString(secretKey)
}
