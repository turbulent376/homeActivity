package auth

import (
	"net/http"

	"git.jetbrains.space/orbi/fcsd/api/internal/public"
	kitHttp "git.jetbrains.space/orbi/fcsd/kit/http"
	pb "git.jetbrains.space/orbi/fcsd/proto/auth"
)

type Controller interface {
	AuthUserByEmail(w http.ResponseWriter, r *http.Request)
	AuthUserByFirebase(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	UserInfo(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	SaveFCMToken(w http.ResponseWriter, r *http.Request)
}

type ctrlImpl struct {
	kitHttp.BaseController
	authRepo public.AuthRepository
}

func NewController(authRepo public.AuthRepository) Controller {
	return &ctrlImpl{
		authRepo: authRepo,
	}
}

// AuthUserByEmail godoc
// @Summary authorization user by email and password
// @Router /auth/login [post]
// @Accept json
// @Produce json
// @Param request body AuthRequest true "request"
// @Success 200 {object} AuthResponse
// @Failure 500 {object} kitHttp.Error
// @tags auth
func (c *ctrlImpl) AuthUserByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var rq AuthRequest

	err := c.DecodeRequest(r, ctx, &rq)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	res, err := c.authRepo.AuthUserByEmail(ctx, c.toAuthRequestPb(&rq))

	if err != nil {
		c.RespondError(w, err)
		return
	}

	c.RespondOK(w, c.toAuthResponse(res))
}

// AuthUserByFirebase godoc
// @Summary authorization user by oauth firebase
// @Router /auth/firebase [post]
// @Accept json
// @Produce json
// @Param request body OAuthRequest true "request"
// @Success 200 {object} AuthResponse
// @Failure 500 {object} kitHttp.Error
// @tags auth
func (c *ctrlImpl) AuthUserByFirebase(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var rq OAuthRequest

	err := c.DecodeRequest(r, ctx, &rq)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	res, err := c.authRepo.AuthUserByFirebase(ctx, &pb.OAuthRequest{
		Token: rq.Token,
	})

	if err != nil {
		c.RespondError(w, err)
		return
	}

	c.RespondOK(w, c.toAuthResponse(res))
}

// RefreshToken godoc
// @Summary refresh token
// @Router /auth/refresh [post]
// @Accept json
// @Produce json
// @Param request body RefreshTokenRequest true "request"
// @Success 200 {object} TokenPairResponse
// @Failure 500 {object} kitHttp.Error
// @tags auth
func (c *ctrlImpl) RefreshToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uid, usid, err := c.CurrentUser(ctx)
	if err != nil {
		c.RespondError(w, err)
		return
	}
	var rq RefreshTokenRequest

	err = c.DecodeRequest(r, ctx, &rq)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	res, err := c.authRepo.RefreshToken(ctx, &pb.RefreshTokenRequest{
		RefreshToken: rq.RefreshToken,
		SessionId:    usid,
		UserId:       uid})

	if err != nil {
		c.RespondError(w, err)
		return
	}

	c.RespondOK(w, c.toTokenPairResponse(res))
}

// CreateUser godoc
// @Summary creating user request
// @Router /auth/user/new [post]
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "request"
// @Success 200 {object} UserResponse
// @Failure 500 {object} kitHttp.Error
// @tags user
func (c *ctrlImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var rq CreateUserRequest

	err := c.DecodeRequest(r, ctx, &rq)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	res, err := c.authRepo.CreateUser(ctx, c.toCreateUserPb(&rq))

	if err != nil {
		c.RespondError(w, err)
		return
	}

	c.RespondOK(w, c.toUserResponse(res))
}

// UserInfo godoc
// @Summary getting user public info
// @Router /auth/user/{userId} [get]
// @Accept json
// @Produce json
// @Success 200 {object} UserInfoResponse
// @Failure 500 {object} kitHttp.Error
// @tags user
func (c *ctrlImpl) UserInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := c.Var(r, ctx, "userId", false)
	if err != nil {
		c.RespondError(w, err)
		return
	}

	res, err := c.authRepo.GetUserById(ctx, &pb.UserIdRequest{Id: id})

	if err != nil {
		c.RespondError(w, err)
		return
	}

	c.RespondOK(w, c.toUserInfoResponse(res))
}

// DeleteUser godoc
// @Summary deleting user
// @Router /auth/user/{userId} [delete]
// @Accept json
// @Produce json
// @Success 200
// @Failure 500 {object} kitHttp.Error
// @tags user
func (c *ctrlImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, _, err := c.CurrentUser(ctx)
	if err != nil {
		c.RespondError(w, err)
		return
	}
	// TODO check permissions

	id, err := c.Var(r, ctx, "userId", false)
	if err != nil {
		c.RespondError(w, err)
		return
	}

	err = c.authRepo.DeleteUser(ctx, &pb.UserIdRequest{Id: id})

	if err != nil {
		c.RespondError(w, err)
		return
	}

	c.RespondOK(w, nil)
}

// UpdateUser godoc
// @Summary updating user
// @Router /auth/user/{userId}} [put]
// @Accept json
// @Produce json
// @Success 200 {object} UserResponse
// @Failure 500 {object} kitHttp.Error
// @tags user
func (c *ctrlImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := c.Var(r, ctx, "userId", false)
	if err != nil {
		c.RespondError(w, err)
		return
	}

	_, _, err = c.CurrentUser(ctx)
	if err != nil {
		c.RespondError(w, err)
		return
	}

	// TODO check persmissions for update user

	var rq UpdateUserRequest
	err = c.DecodeRequest(r, ctx, &rq)
	if err != nil {
		c.RespondError(w, err)
		return
	}

	res, err := c.authRepo.UpdateUser(ctx, c.toUpdateUserRequestPb(&UpdateUserRequest{
		Id:          id,
		Name:        rq.Name,
		Avatar:      rq.Avatar,
		Surname:     rq.Surname,
		Email:       rq.Email,
		CountryCode: rq.CountryCode,
		OldPassword: rq.OldPassword,
		NewPassword: rq.NewPassword,
	}))

	if err != nil {
		c.RespondError(w, err)
		return
	}

	c.RespondOK(w, c.toUserResponse(res))
}

// Logout godoc
// @Summary close current session
// @Router /auth/logout [get]
// @Accept json
// @Produce json
// @Success 200
// @Failure 500 {object} kitHttp.Error
// @tags user
func (c *ctrlImpl) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uid, usid, err := c.CurrentUser(ctx)
	if err != nil {
		c.RespondError(w, err)
		return
	}

	err = c.authRepo.CloseSession(ctx, &pb.CloseSessionRequest{UserId: uid,
		SessionId: usid})

	if err != nil {
		c.RespondError(w, err)
		return
	}

	c.RespondOK(w, nil)
}

// SaveFCMToken godoc
// @Summary save new fcm token for user
// @Router /auth/user/notify/token [post]
// @Accept json
// @Produce json
// @Param request body FCMTokenRequest true "request"
// @Success 200
// @Failure 500 {object} kitHttp.Error
// @tags user
func (c *ctrlImpl) SaveFCMToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uid, usid, err := c.CurrentUser(ctx)
	if err != nil {
		c.RespondError(w, err)
		return
	}

	var rq FCMTokenRequest

	err = c.DecodeRequest(r, ctx, &rq)
	if err != nil {
		c.RespondError(w, err)
		return
	}

	err = c.authRepo.SaveUserFCMToken(ctx, &pb.FCMTokenRequest{
		UserId:    uid,
		SessionId: usid,
		FCMToken:  rq.FCMToken})

	if err != nil {
		c.RespondError(w, err)
		return
	}

	c.RespondOK(w, nil)
}
