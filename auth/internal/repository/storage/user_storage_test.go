//+test integration

package storage

import (
	"context"
	"testing"
	"time"

	"git.jetbrains.space/orbi/fcsd/auth/internal/config"
	"git.jetbrains.space/orbi/fcsd/auth/internal/domain"
	kitContext "git.jetbrains.space/orbi/fcsd/kit/context"
	"git.jetbrains.space/orbi/fcsd/kit/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type authStorageTestSuite struct {
	suite.Suite
	storage     domain.UserStorage
	ctx         context.Context
	defaultUser *domain.User
}

// SetupSuite is called once for a suite
func (s *authStorageTestSuite) SetupSuite() {

	// setup context
	s.ctx = kitContext.NewRequestCtx().Test().ToContext(context.Background())

	// load config
	cfg, err := config.Load()
	if err != nil {
		s.T().Fatal(err)
	}

	// disable applying migrations
	cfg.Storages.Database.MigPath = ""

	// initialize adapter
	a := NewAdapter()
	err = a.Init(cfg.Storages)
	if err != nil {
		s.T().Fatal(err)
	}
	s.storage = a.GetUserStorage()
}

// SetupTest is called for each test
func (s *authStorageTestSuite) SetupTest() {
	now := time.Now().UTC()
	id := utils.NewId()
	s.defaultUser = &domain.User{
		Id:           id,
		Name:         "name",
		Surname:      "surname",
		Email:        "test@test.test",
		FirebaseUUID: "UUID",
		KundelikId:   "011100",
		CreatedAt:    now,
		UpdatedAt:    now,
		DeletedAt:    nil,
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(authStorageTestSuite))
}

func (s *authStorageTestSuite) Test_CreateUser_GetFromDbAndCache() {

	// create a task
	expected := s.defaultUser
	err := s.storage.CreateUser(s.ctx, expected)
	if err != nil {
		s.T().Fatal(err)
	}

	// get User by id
	actual, err := s.storage.GetUser(s.ctx, expected.Id)
	if actual == nil || err != nil {
		s.T().Fatal()
	}

	assert.Equal(s.T(), expected.Id, actual.Id)
	assert.Equal(s.T(), expected.Name, actual.Name)
	assert.Equal(s.T(), expected.Surname, actual.Surname)
	assert.Equal(s.T(), expected.Email, actual.Email)

	// wait for async caching
	time.Sleep(time.Millisecond * 100)

	// get Task by id again (cache hit)
	actual, err = s.storage.GetUser(s.ctx, expected.Id)
	if actual == nil || err != nil {
		s.T().Fatal()
	}
	assert.Equal(s.T(), expected.Id, actual.Id)
	assert.Equal(s.T(), expected.Name, actual.Name)
	assert.Equal(s.T(), expected.Surname, actual.Surname)
	assert.Equal(s.T(), expected.Email, actual.Email)
}

func (s *authStorageTestSuite) TestUpdateUser() {

	// create a sample
	expected := s.defaultUser
	err := s.storage.CreateUser(s.ctx, expected)
	if err != nil {
		s.T().Fatal(err)
	}

	// get Task by id
	actual, err := s.storage.GetUser(s.ctx, expected.Id)
	if err != nil {
		s.T().Fatal()
	}
	assert.NotEmpty(s.T(), actual.Id)

	// update user
	actual.Name = "another name"
	err = s.storage.UpdateUser(s.ctx, actual)
	if err != nil {
		s.T().Fatal(err)
	}

	// get Sample by id
	actual, err = s.storage.GetUser(s.ctx, expected.Id)
	if err != nil {
		s.T().Fatal()
	}
	assert.NotEmpty(s.T(), actual.Id)
	assert.Equal(s.T(), actual.Name, "another name")
}

func (s *authStorageTestSuite) Test_GetUserByEmail() {

	// create a task
	expected := s.defaultUser
	expected.Email = "new-email@email.email"
	err := s.storage.CreateUser(s.ctx, expected)
	if err != nil {
		s.T().Fatal(err)
	}
	// get Task by email
	actual, err := s.storage.GetUserByEmail(s.ctx, expected.Email)
	if actual == nil || err != nil {
		s.T().Fatal()
	}
	assert.Equal(s.T(), expected.Id, actual.Id)
	assert.Equal(s.T(), expected.Name, actual.Name)
	assert.Equal(s.T(), expected.Surname, actual.Surname)
	assert.Equal(s.T(), expected.Email, actual.Email)

	// wait for async caching
	time.Sleep(time.Millisecond * 100)

	// get Task by email
	actual, err = s.storage.GetUserByEmail(s.ctx, expected.Email)
	if actual == nil || err != nil {
		s.T().Fatal()
	}
	assert.Equal(s.T(), expected.Id, actual.Id)
	assert.Equal(s.T(), expected.Name, actual.Name)
	assert.Equal(s.T(), expected.Surname, actual.Surname)
	assert.Equal(s.T(), expected.Email, actual.Email)

}
