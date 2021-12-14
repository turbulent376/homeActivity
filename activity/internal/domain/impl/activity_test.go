package impl

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/turbulent376/homeactivity/activity/internal/domain"
	"github.com/turbulent376/homeactivity/activity/internal/mocks"
	kitContext "github.com/turbulent376/kit/context"
	kitTest "github.com/turbulent376/kit/test"
	kitUtils "github.com/turbulent376/kit/utils"
	pb "github.com/turbulent376/proto/activity"
	"testing"
	"time"
)

type activityTestSuite struct {
	suite.Suite
	activityStorage     *mocks.ActivityStorage
	activityTypeStorage *mocks.ActivityTypeStorage
	activityService     domain.ActivityService
	ctx              context.Context
}

func (s *activityTestSuite) SetupSuite() {
	s.activityStorage = &mocks.ActivityStorage{}
	s.activityTypeStorage = &mocks.ActivityTypeStorage{}
	s.ctx = kitContext.NewRequestCtx().Test().ToContext(context.Background())
	s.activityService = NewActivityService(s.activityStorage, s.activityTypeStorage)
}

func (s *activityTestSuite) SetupTest() {
	s.activityStorage.ExpectedCalls = nil
	s.activityStorage.On("CreateTimetable",
		mock.AnythingOfType("*context.valueCtx"),
		mock.AnythingOfType("*domain.Timesheet")).
		Return(&domain.Activity{
			Id:        kitUtils.NewId(),
			Owner:     "123",
			DateFrom:  time.Now(),
			DateTo:    time.Now(),
		}, nil)
	s.activityStorage.On("UpdateTimetable",
		mock.AnythingOfType("*context.valueCtx"),
		mock.AnythingOfType("*domain.Timesheet")).
		Return(&domain.Activity{
			Id:        kitUtils.NewId(),
			Owner:     "123",
			DateFrom:  time.Now(),
			DateTo:    time.Now(),
		}, nil)
	s.activityStorage.On("GetTimetable",
		mock.AnythingOfType("*context.valueCtx"),
		mock.AnythingOfType("string")).
		Return(true, &domain.Activity{
			Id:        kitUtils.NewId(),
			Owner:     "123",
			DateFrom:  time.Now(),
			DateTo:    time.Now(),
		}, nil)
	s.activityStorage.On("SearchTimetable",
		mock.AnythingOfType("*context.valueCtx"),
		mock.AnythingOfType("*domain.SearchTimesheetRequest")).
		Return(true, &domain.Activity{
			Id:        kitUtils.NewId(),
			Owner:     "123",
			DateFrom:  time.Time{},
			DateTo:    time.Now(),
		}, nil)
	s.activityStorage.On("DeleteTimetable",
		mock.AnythingOfType("*context.valueCtx"),
		mock.AnythingOfType("string")).
		Return(nil)

	s.activityTypeStorage.ExpectedCalls = nil
	s.activityTypeStorage.On("CreateEvent",
		mock.AnythingOfType("*context.valueCtx"),
		mock.AnythingOfType("*domain.Event")).
		Return(&domain.ActivityType{
			Id:          kitUtils.NewId(),
		}, nil)
	s.activityTypeStorage.On("UpdateEvent",
		mock.AnythingOfType("*context.valueCtx"),
		mock.AnythingOfType("*domain.Event")).
		Return(&domain.ActivityType{
			Id:          kitUtils.NewId(),
		}, nil)
	s.activityTypeStorage.On("GetEvent",
		mock.AnythingOfType("*context.valueCtx"),
		mock.AnythingOfType("string")).
		Return(true, &domain.ActivityType{
			Id:          kitUtils.NewId(),
		}, nil)
	s.activityTypeStorage.On("DeleteEvent",
		mock.AnythingOfType("*context.valueCtx"),
		mock.AnythingOfType("string")).
		Return(nil)
}

func TestActivitySuite(t *testing.T) {
	suite.Run(t, new(activityTestSuite))
}

func (s *activityTestSuite) Test_Create_WhenEmptyName_Fail() {
	_, err := s.activityService.Create(s.ctx, &domain.Activity{})
	kitTest.AssertAppErr(s.T(), err, pb.ErrCodeActivityOwnerEmpty)
}

func (s *activityTestSuite) Test_Create_Ok() {

	tt, err := s.activityService.Create(s.ctx, &domain.Activity{
		Owner:    "123",
		DateFrom: time.Now(),
		DateTo:   time.Now(),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	assert.NotEmpty(s.T(), tt)
	assert.NotEmpty(s.T(), tt.Id)
	assert.NotEmpty(s.T(), tt.Owner)
	assert.NotEmpty(s.T(), tt.DateFrom)
	assert.NotEmpty(s.T(), tt.DateTo)
}

func (s *activityTestSuite) Test_Update_Fail() {
	_, err := s.activityService.Update(s.ctx, &domain.Activity{})
	kitTest.AssertAppErr(s.T(), err, pb.ErrCodeActivityIdEmpty)
}

func (s *activityTestSuite) Test_Update_Ok() {
	_, err := s.activityService.Update(s.ctx, &domain.Activity{
		Id:       kitUtils.NewId(),
		Owner:    "123",
		DateFrom: time.Now(),
		DateTo:   time.Now(),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	assert.Nil(s.T(), err)

}

func (s *activityTestSuite) Test_Get_Fail() {
	_, _, err := s.activityService.Get(s.ctx, "")
	kitTest.AssertAppErr(s.T(), err, pb.ErrCodeActivityIdEmpty)
}

func (s *activityTestSuite) Test_Get_Ok() {
	_, tt, err := s.activityService.Get(s.ctx, kitUtils.NewId())
	if err != nil {
		s.T().Fatal(err)
	}
	assert.Nil(s.T(), err)
	assert.NotEmpty(s.T(), tt)
}


func (s *activityTestSuite) Test_Delete_Ok() {
	err := s.activityService.Delete(s.ctx, kitUtils.NewId())
	if err != nil {
		s.T().Fatal(err)
	}
	assert.Nil(s.T(), err)
}

func (s *activityTestSuite) Test_CreateActivityType_WhenEmptyName_Fail() {
	_, err := s.activityService.CreateActivityType(s.ctx, &domain.ActivityType{})
	kitTest.AssertAppErr(s.T(), err, pb.ErrCodeActivityNameIsEmpty)
}

func (s *activityTestSuite) Test_CreateActivityType_Ok() {

	ev, err := s.activityService.CreateActivityType(s.ctx, &domain.ActivityType{
		Id:          "321",
		Family: "321",
		Name:     "test subject",
		Description:     "Monday",
	})
	if err != nil {
		s.T().Fatal(err)
	}
	assert.NotEmpty(s.T(), ev)
	assert.NotEmpty(s.T(), ev.Id)
	assert.NotEmpty(s.T(), ev.Family)
	assert.NotEmpty(s.T(), ev.Name)
	assert.NotEmpty(s.T(), ev.Description)
}

func (s *activityTestSuite) Test_UpdateActivityType_Fail() {
	_, err := s.activityService.UpdateActivityType(s.ctx, &domain.ActivityType{})
	kitTest.AssertAppErr(s.T(), err, pb.ErrCodeActivityNotFound)
}

func (s *activityTestSuite) Test_UpdateEvent_Ok() {
	_, err := s.activityService.UpdateActivityType(s.ctx, &domain.ActivityType{
		Id:          kitUtils.NewId(),
		Family: kitUtils.NewId(),
		Name:     "test subject 2",
		Description:     "Monday",
	})
	if err != nil {
		s.T().Fatal(err)
	}
	assert.Nil(s.T(), err)
}

func (s *activityTestSuite) Test_GetEvent_Fail() {
	_, tt, err := s.activityService.GetActivityType(s.ctx, "123")
	if err != nil {
		s.T().Fatal(err)
	}
	assert.Nil(s.T(), err)
	assert.NotEmpty(s.T(), tt)
	assert.NotEqual(s.T(), tt, &domain.ActivityType{})
}

func (s *activityTestSuite) Test_GetActivityType_Ok() {
	_, tt, err := s.activityService.GetActivityType(s.ctx, "321")
	if err != nil {
		s.T().Fatal(err)
	}
	assert.Nil(s.T(), err)
	assert.NotEmpty(s.T(), tt)
}

func (s *activityTestSuite) Test_DeleteActivityType_Ok() {
	err := s.activityService.DeleteActivityType(s.ctx, "321")
	if err != nil {
		s.T().Fatal(err)
	}
	assert.Nil(s.T(), err)
}
