package impl

import (
	"context"
	"testing"

	"git.jetbrains.space/orbi/fcsd/auth/internal/domain"
	"git.jetbrains.space/orbi/fcsd/auth/internal/mocks"
	kitContext "git.jetbrains.space/orbi/fcsd/kit/context"
	kitTest "git.jetbrains.space/orbi/fcsd/kit/test"
	kitUtils "git.jetbrains.space/orbi/fcsd/kit/utils"
	pb "git.jetbrains.space/orbi/fcsd/proto/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type authTestSuite struct {
	suite.Suite
	userStorage         *mocks.UserStorage
	sessionStorage      *mocks.SessionStorage
	notificationAdapter *mocks.NotificationAdapter
	firebaseAdapter     *mocks.FirebaseAdapter
	authService         domain.AuthService
	ctx                 context.Context
}

func (a *authTestSuite) SetupSuite() {
	a.userStorage = &mocks.UserStorage{}
	a.sessionStorage = &mocks.SessionStorage{}
	a.notificationAdapter = &mocks.NotificationAdapter{}
	a.firebaseAdapter = &mocks.FirebaseAdapter{}
	a.ctx = kitContext.NewRequestCtx().Test().ToContext(context.Background())
}

func (a *authTestSuite) SetupTest() {
	a.userStorage.ExpectedCalls = nil
	a.userStorage.On("CreateUser", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("*domain.User")).Return(nil)
	a.userStorage.On("UpdateUser", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("*domain.User")).Return(nil)
	a.userStorage.On("DeleteUser", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string")).Return(nil)
	a.userStorage.
		On("GetUser", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string")).
		Return(&domain.User{
			Id:       kitUtils.NewId(),
			Email:    "email@email.ru",
			Password: "$2a$10$AToD3K98R.jrQGvEVYWnU.nA3Evt0u.0kksHV1F1jl4JocW/EkPW6",
			Name:     "1",
		}, nil)
	a.sessionStorage.On("GetUserSessions", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string")).Return([]*domain.Session{
		{
			Id:           kitUtils.NewId(),
			UserId:       "111",
			DeviceName:   "Nexus Android 10",
			RefreshToken: "asdfg",
			FCMToken:     "",
		},
		{
			Id:           kitUtils.NewId(),
			UserId:       "111",
			DeviceName:   "iPhone 13 pro",
			RefreshToken: "qwerty",
			FCMToken:     "token1",
		},
		{
			Id:           kitUtils.NewId(),
			UserId:       "111",
			DeviceName:   "Nokia3110",
			RefreshToken: "qwerty",
			FCMToken:     "token2",
		},
	}, nil)
	a.notificationAdapter.On("SendNotify", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string")).Return(nil, nil)
	a.authService = NewAuthService(a.userStorage, a.sessionStorage, a.notificationAdapter, a.firebaseAdapter)
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(authTestSuite))
}

func (a *authTestSuite) Test_CreateUser_WhenEmptyEmail_Fail() {
	_, err := a.authService.CreateUser(a.ctx, &domain.RegistrationUserRequest{
		Password: "123456",
	})
	kitTest.AssertAppErr(a.T(), err, pb.ErrCodeAuthInvalidEmail)
}

func (a *authTestSuite) Test_CreateUser_WhenEmptyPassword_Fail() {
	a.userStorage.On("GetUserByEmail", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string")).Return(nil, nil)
	_, err := a.authService.CreateUser(a.ctx, &domain.RegistrationUserRequest{
		Email: "email@email.ru",
	})
	kitTest.AssertAppErr(a.T(), err, pb.ErrCodeAuthInvalidPassword)
}

func (a *authTestSuite) Test_CreateUser_Ok() {
	a.userStorage.On("GetUserByEmail", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string")).Return(nil, nil)
	user, err := a.authService.CreateUser(a.ctx, &domain.RegistrationUserRequest{
		Email:    "user@email.ru",
		Password: "12345678",
	})
	if err != nil {
		a.T().Fatal(err)
	}
	assert.Nil(a.T(), err)
	assert.NotEmpty(a.T(), user.Id)

}

func (a *authTestSuite) Test_CreateUser_EmailAlreadyExist_fail() {
	a.userStorage.
		On("GetUserByEmail", mock.AnythingOfType("*context.valueCtx"), mock.AnythingOfType("string")).
		Return(&domain.User{
			Id:    kitUtils.NewId(),
			Email: "email@email.ru",
		}, nil)
	_, err := a.authService.CreateUser(a.ctx, &domain.RegistrationUserRequest{
		Email:    "email@email.ru",
		Password: "12345678",
	})
	kitTest.AssertAppErr(a.T(), err, pb.ErrCodeAuthInvalidEmail)

}

func (a *authTestSuite) Test_DeleteUser_Ok() {
	err := a.authService.DeleteUser(a.ctx, "111")
	assert.Nil(a.T(), err)
}

func (a *authTestSuite) Test_UpdateUser_WhenEmptyEmail_Fail() {
	_, err := a.authService.UpdateUser(a.ctx, &domain.UpdateUserRequest{
		Id:      "444",
		Name:    "Name",
		Surname: "Surname",
	})
	kitTest.AssertAppErr(a.T(), err, pb.ErrCodeAuthInvalidEmail)
}
func (a *authTestSuite) Test_UpdateUser_Ok() {
	_, err := a.authService.UpdateUser(a.ctx, &domain.UpdateUserRequest{
		Id:      "444",
		Email:   "new@email.ru",
		Name:    "Name",
		Surname: "Surname",
	})
	assert.Nil(a.T(), err)
}

func (a *authTestSuite) Test_UpdateUser_BadOldPass_Fail() {
	_, err := a.authService.UpdateUser(a.ctx, &domain.UpdateUserRequest{
		Id:          "444",
		Email:       "new@email.ru",
		Name:        "Name",
		Surname:     "Surname",
		OldPassword: "00000000",
		NewPassword: "11111111",
	})
	kitTest.AssertAppErr(a.T(), err, pb.ErrCodeAuthInvalidPassword)
}

func (a *authTestSuite) Test_UpdateUser_NewPass_Ok() {
	_, err := a.authService.UpdateUser(a.ctx, &domain.UpdateUserRequest{
		Id:          "444",
		Email:       "new@email.ru",
		Name:        "Name",
		Surname:     "Surname",
		OldPassword: "12345678",
		NewPassword: "11111111",
	})
	assert.Nil(a.T(), err)
}

func (a *authTestSuite) Test_GetUserFCMTokens_WithoutEmptyString_Ok() {
	tokens, err := a.authService.GetUserFCMTokens(a.ctx, "1")
	assert.Nil(a.T(), err)
	assert.Equal(a.T(), len(tokens), 2)
}

func (a *authTestSuite) Test_GetUserFCMTokens_Ok() {
	tokens, err := a.authService.GetUserFCMTokens(a.ctx, "1")
	assert.Nil(a.T(), err)
	assert.Equal(a.T(), tokens, []string{"token1", "token2"})
}
