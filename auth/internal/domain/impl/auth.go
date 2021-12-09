package impl

import (
	"context"
	"crypto/rsa"
	"time"

	"git.jetbrains.space/orbi/fcsd/auth/internal/domain"
	"git.jetbrains.space/orbi/fcsd/auth/internal/errors"
	"git.jetbrains.space/orbi/fcsd/auth/internal/logger"
	"git.jetbrains.space/orbi/fcsd/auth/internal/repository/adapters/firebase"
	"git.jetbrains.space/orbi/fcsd/auth/internal/repository/adapters/notification"
	"git.jetbrains.space/orbi/fcsd/kit/log"
	"git.jetbrains.space/orbi/fcsd/kit/utils"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authImpl struct {
	userStorage         domain.UserStorage
	sessionStorage      domain.SessionStorage
	notificationAdapter notification.NotificationAdapter
	firebaseAdapter     firebase.FirebaseAdapter
	signKey             *rsa.PrivateKey
}

func NewAuthService(userStorage domain.UserStorage,
	sessionStorage domain.SessionStorage,
	notificationAdapter notification.NotificationAdapter,
	firebaseAdapter firebase.FirebaseAdapter,
) domain.AuthService {
	return &authImpl{
		userStorage:         userStorage,
		sessionStorage:      sessionStorage,
		notificationAdapter: notificationAdapter,
		firebaseAdapter:     firebaseAdapter,
	}
}
func (*authImpl) l() log.CLogger {
	return logger.L().Cmp("auth")
}

func (a *authImpl) privateKey() *rsa.PrivateKey {
	return a.signKey
}
func (a *authImpl) publicKey() *rsa.PublicKey {
	return &a.signKey.PublicKey
}

func (a *authImpl) genTokenPair(ctx context.Context, userId, sessionId string) (*domain.TokenPair, error) {
	tk := &domain.Token{
		AppId:          domain.AuId,
		UserId:         userId,
		SessionId:      sessionId,
		StandardClaims: jwt.StandardClaims{},
	}
	tk.StandardClaims.ExpiresAt = time.Now().Add(domain.TokenExp).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), tk)

	tokenString, err := token.SignedString(a.privateKey())
	if err != nil {
		return nil, errors.ErrAuthSignToken(err, ctx)
	}

	rt := domain.RefreshToken{
		UserId:         userId,
		StandardClaims: jwt.StandardClaims{},
	}
	rt.StandardClaims.ExpiresAt = time.Now().Add(domain.RefreshTokenExp).Unix()
	refreshToken := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), rt)

	newRefreshToken, err := refreshToken.SignedString(a.privateKey())
	if err != nil {
		return nil, errors.ErrAuthSignToken(err, ctx)
	}

	tokPair := &domain.TokenPair{
		Token:        tokenString,
		RefreshToken: newRefreshToken,
	}
	return tokPair, nil
}

func (a *authImpl) newSession(ctx context.Context, userId string) (*domain.TokenPair, error) {

	session := &domain.Session{
		Id:        utils.NewId(),
		UserId:    userId,
		CreatedAt: time.Now().UTC(),
	}

	tokPair, err := a.genTokenPair(ctx, userId, session.Id)
	if err != nil {
		return nil, errors.ErrAuthSignToken(err, ctx)
	}
	session.RefreshToken = tokPair.RefreshToken

	err = a.sessionStorage.CreateSession(ctx, session)
	if err != nil {
		return nil, errors.ErrSessionCreate(ctx, session.Id)
	}
	return tokPair, nil

}

func (a *authImpl) SetSignKey(ctx context.Context, key string) error {
	var err error
	a.signKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(key))
	if err != nil {
		return errors.ErrAuthSignToken(err, ctx)
	}
	return nil
}

func (a *authImpl) AuthUserByEmail(ctx context.Context, req *domain.AuthRequest) (*domain.AuthResponse, error) {
	a.l().C(ctx).Mth("AuthUserByEmail").Dbg()
	user, err := a.userStorage.GetUserByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return nil, errors.ErrGetUserByField(ctx, "email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errors.ErrUserInvalidPassword(ctx)
	}

	tokenPair, err := a.newSession(ctx, user.Id)
	if err != nil {
		return nil, errors.ErrSessionCreate(ctx, user.Id)
	}
	return &domain.AuthResponse{
		Token:        tokenPair.Token,
		RefreshToken: tokenPair.RefreshToken,
		User:         *user,
	}, nil
}

func (a *authImpl) AuthUserByFirebase(ctx context.Context, req *domain.OAuthRequest) (*domain.AuthResponse, error) {
	a.l().C(ctx).Mth("AuthUserByFirebase").Dbg()
	if req.Token == "" {
		return nil, errors.ErrAuthEmptyToken(ctx)
	}
	authToken, err := a.firebaseAdapter.VerifyIDToken(ctx, req.Token)
	if err != nil {
		return nil, errors.ErrFirebaseVerifyToken(err, ctx)
	}
	user, err := a.userStorage.GetUserByFireUUID(ctx, authToken.UID)
	if err != nil {
		return nil, errors.ErrGetUserByField(ctx, "firebaseUUID")
	}
	if user == nil {
		user = &domain.User{}
		now := time.Now().UTC()
		user.Id = utils.NewId()
		user.FirebaseUUID = authToken.UID
		user.CreatedAt, user.UpdatedAt = now, now

		// save to store
		err = a.userStorage.CreateUser(ctx, user)
		if err != nil {
			return nil, errors.ErrUserCreate(ctx, user.Id)
		}
	}

	tokenPair, err := a.newSession(ctx, user.Id)
	if err != nil {
		return nil, errors.ErrSessionCreate(ctx, user.Id)
	}
	return &domain.AuthResponse{
		Token:        tokenPair.Token,
		RefreshToken: tokenPair.RefreshToken,
		User:         *user,
	}, nil

}

func (a *authImpl) RefreshToken(ctx context.Context, req *domain.RefreshTokenRequest) (*domain.TokenPair, error) {
	a.l().C(ctx).Mth("RefreshToken").Dbg()

	session, err := a.sessionStorage.GetSession(ctx, req.SessionId)
	if err != nil {
		return nil, errors.ErrSessionNotFound(err, ctx)
	}
	if session == nil {
		return nil, errors.ErrUserEmptySession(ctx, req.UserId)
	}

	if session.RefreshToken != req.RefreshToken || session.UserId != req.UserId {
		return nil, errors.ErrAuthInvalidRefreshToken(ctx)
	}

	rtok, err := jwt.ParseWithClaims(req.RefreshToken, &domain.RefreshToken{}, func(token *jwt.Token) (interface{}, error) {
		return a.publicKey(), nil
	})

	if err != nil {
		return nil, errors.ErrAuthRefreshToken(err, ctx)
	}
	if _, ok := rtok.Claims.(*domain.RefreshToken); !ok || !rtok.Valid {
		return nil, errors.ErrAuthInvalidRefreshToken(ctx)
	}

	tokPair, err := a.genTokenPair(ctx, req.UserId, req.SessionId)
	if err != nil {
		return nil, errors.ErrAuthRefreshToken(err, ctx)
	}

	return tokPair, nil
}
func (a *authImpl) CreateUser(ctx context.Context, rq *domain.RegistrationUserRequest) (*domain.User, error) {
	a.l().C(ctx).Mth("CreateUser").Dbg()

	// validates name
	if rq.Email == "" {
		return nil, errors.ErrUserInvalidEmail(ctx)
	}

	stored, err := a.userStorage.GetUserByEmail(ctx, rq.Email)
	if err != nil {
		return nil, errors.ErrGetUserByField(ctx, "email")
	}
	if stored != nil {
		return nil, errors.ErrEmailAlreadyExist(ctx)
	}

	if rq.Password == "" {
		return nil, errors.ErrUserInvalidPassword(ctx)
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(rq.Password), bcrypt.DefaultCost)
	user := &domain.User{
		Password: string(hashedPassword),
		Email:    rq.Email,
		Name:     rq.Name,
		Surname:  rq.Surname,
	}
	now := time.Now().UTC()
	user.Id = utils.NewId()
	user.CreatedAt, user.UpdatedAt = now, now

	// save to store
	err = a.userStorage.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.ErrUserCreate(ctx, user.Id)
	}

	return user, nil
}

func (a *authImpl) UpdateUser(ctx context.Context, rq *domain.UpdateUserRequest) (*domain.User, error) {
	a.l().C(ctx).Mth("UpdateUser").Dbg()

	// retrieve stored user by id
	stored, err := a.userStorage.GetUser(ctx, rq.Id)
	if err != nil {
		return nil, errors.ErrUserNotFound(ctx, rq.Id)
	}
	if stored == nil {
		return nil, errors.ErrUserNotFound(ctx, rq.Id)
	}

	if rq.NewPassword != "" && len(rq.NewPassword) > 6 {
		if rq.OldPassword != "" && len(rq.OldPassword) > 6 {
			err = bcrypt.CompareHashAndPassword([]byte(stored.Password), []byte(rq.OldPassword))
			if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
				return nil, errors.ErrUserInvalidPassword(ctx)
			}
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(rq.NewPassword), bcrypt.DefaultCost)
			stored.Password = string(hashedPassword)

		}
	}
	if rq.Email == "" {
		return nil, errors.ErrUserInvalidEmail(ctx)
	}

	stored.Email = rq.Email
	stored.Avatar = rq.Avatar
	stored.Name = rq.Name
	stored.Surname = rq.Surname
	stored.CountryCode = rq.CountryCode

	// set updated params
	now := time.Now().UTC()
	stored.UpdatedAt = now

	// save to store
	err = a.userStorage.UpdateUser(ctx, stored)
	if err != nil {
		return nil, errors.ErrUserUpdate(ctx, rq.Id)
	}

	fmcTokens, err := a.GetUserFCMTokens(ctx, stored.Id)
	if err != nil {
		return nil, errors.ErrUserSendNotification(err, ctx)
	}
	for _, t := range fmcTokens {
		_, err = a.notificationAdapter.SendNotify(ctx, t)
		if err != nil {
			return nil, errors.ErrUserSendNotification(err, ctx)
		}
	}

	return stored, nil
}

func (a *authImpl) GetUserById(ctx context.Context, id string) (*domain.User, error) {
	a.l().C(ctx).Mth("GetUserById").Dbg()
	return a.userStorage.GetUser(ctx, id)
}

func (a *authImpl) GetSessionByToken(ctx context.Context, token string) (*domain.Session, error) {
	a.l().C(ctx).Mth("GetSessionByToken").Dbg()

	tok, err := jwt.ParseWithClaims(token, &domain.Token{}, func(t *jwt.Token) (interface{}, error) {
		return a.publicKey(), nil
	})
	if err != nil {
		return nil, errors.ErrParseJWT(err, ctx)
	}

	tokClaims, ok := tok.Claims.(*domain.Token)
	if !ok || !tok.Valid {
		return nil, errors.ErrAuthInvalidRefreshToken(ctx)
	}

	return a.sessionStorage.GetSession(ctx, tokClaims.SessionId)
}

func (a *authImpl) DeleteUser(ctx context.Context, id string) error {
	a.l().C(ctx).Mth("DeleteUser").Dbg()

	// check Id isn't empty
	if id == "" {
		return errors.ErrUserIdEmpty(ctx)
	}

	// retrieve stored user by id
	stored, err := a.userStorage.GetUser(ctx, id)
	if err != nil {
		return errors.ErrGetUserByField(ctx, "id")
	}
	if stored == nil {
		return errors.ErrUserNotFound(ctx, id)
	}

	// check already deleted
	if stored.DeletedAt != nil {
		return errors.ErrUserDelete(ctx, id)
	}

	// set updated params
	now := time.Now().UTC()
	stored.UpdatedAt, stored.DeletedAt = now, &now

	// save to store
	return a.userStorage.UpdateUser(ctx, stored)
}

func (a *authImpl) CloseSession(ctx context.Context, req *domain.CloseSessionRequest) error {
	a.l().C(ctx).Mth("CloseSession").Dbg()
	err := a.sessionStorage.DeleteSession(ctx, req.SessionId)
	if err != nil {
		return errors.ErrSessionDelete(ctx, req.SessionId)
	}
	return nil

}

func (a *authImpl) GetUserFCMTokens(ctx context.Context, id string) ([]string, error) {
	a.l().C(ctx).Mth("GetUserFCMTokens").Dbg()
	sessions, err := a.sessionStorage.GetUserSessions(ctx, id)
	if err != nil {
		return nil, errors.ErrSessionNotFound(err, ctx)
	}
	if sessions == nil {
		return nil, errors.ErrUserEmptySession(ctx, id)
	}

	var fmcTokens []string
	for _, value := range sessions {
		if value.FCMToken != "" {
			fmcTokens = append(fmcTokens, value.FCMToken)
		}
	}
	return fmcTokens, nil
}

func (a *authImpl) SaveUserFCMToken(ctx context.Context, req *domain.FCMTokenRequest) error {
	a.l().C(ctx).Mth("SaveUserFCMToken").Dbg()
	session, err := a.sessionStorage.GetSession(ctx, req.SessionId)
	if err != nil {
		return errors.ErrSessionNotFound(err, ctx)
	}
	if session == nil {
		return errors.ErrUserEmptySession(ctx, req.UserId)
	}

	session.FCMToken = req.FCMToken
	err = a.sessionStorage.UpdateSession(ctx, session)
	if err != nil {
		return errors.ErrSessionUpdate(ctx, session.Id)
	}
	return nil
}

func (a *authImpl) GetUserSessions(ctx context.Context, id string) ([]*domain.Session, error) {
	a.l().C(ctx).Mth("GetUserSessions").Dbg()
	sessions, err := a.sessionStorage.GetUserSessions(ctx, id)
	if err != nil {
		return nil, errors.ErrSessionNotFound(err, ctx)
	}
	if sessions == nil {
		return nil, errors.ErrUserEmptySession(ctx, id)
	}

	return sessions, nil
}

func (a *authImpl) SetSessionInfo(ctx context.Context, req *domain.SetSessionInfoRequest) error {
	a.l().C(ctx).Mth("SetSessionInfo").Dbg()
	session, err := a.sessionStorage.GetSession(ctx, req.SessionId)
	if err != nil || session == nil {
		return errors.ErrSessionNotFound(err, ctx)
	}

	session.DeviceName = req.DeviceName
	session.ClientVersion = req.ClientVersion
	err = a.sessionStorage.UpdateSession(ctx, session)
	if err != nil {
		return errors.ErrSessionUpdate(ctx, session.Id)
	}
	return nil

}
