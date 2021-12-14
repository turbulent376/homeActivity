//go:build integration
// +build integration

package grpc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/turbulent376/homeactivity/activity/internal/config"
	kitContext "github.com/turbulent376/kit/context"
	kitGrpc "github.com/turbulent376/kit/grpc"
	pb "github.com/turbulent376/proto/activity"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

type activityGrpcTestSuite struct {
	suite.Suite
	ctx         context.Context
	clActivity pb.ActivityServiceClient
}

func (s *activityGrpcTestSuite) SetupSuite() {

	// setup context
	s.ctx = kitContext.NewRequestCtx().Test().ToContext(context.Background())

	// load config
	cfg, err := config.Load()
	if err != nil {
		s.T().Fatal(err)
	}

	// create GRPC client
	cl, err := kitGrpc.NewClient(&kitGrpc.ClientConfig{Host: cfg.Grpc.Host, Port: cfg.Grpc.Port})
	if err != nil {
		s.T().Fatal(err)
	}
	s.clAtivity = pb.NewActivityServiceClient(cl.Conn)
}

func TestActivitySuite(t *testing.T) {
	suite.Run(t, new(activityGrpcTestSuite))
}

func (s *timesheetGrpcTestSuite) getCreateActivityRequest() *pb.CreateTimesheetRequest {

	return &pb.CreateActivityRequest{
	}
}

func (s *activityGrpcTestSuite) getCreateActivityTypeRequest() *pb.CreateActivityTypeRequest {

	return &pb.CreateActivityTypeRequest{
	}
}

func (s *activityGrpcTestSuite) TestActivityCRUD() {

	// create a new consultant
	cl, err := s.clActivity.Create(s.ctx, s.getCreateActivityRequest())
	if err != nil {
		s.T().Fatal(err)
	}
	assert.NotEmpty(s.T(), cl.Id)

	// get by id
	cl, err = s.clActivity.Get(s.ctx, &pb.ActivityIdRequest{Id: cl.Id})
	if err != nil {
		s.T().Fatal()
	}
	assert.NotEmpty(s.T(), cl)
	assert.NotEmpty(s.T(), cl.Id)

	// set status to active
	cl, err = s.clActivity.Update(s.ctx, &pb.UpdateActivityRequest{
		Id:       cl.Id,
		Owner:    cl.Owner,
		DateFrom: timestamppb.Now(),
		DateTo:   timestamppb.Now(),
	})
	if err != nil {
		s.T().Fatal()
	}
	assert.Equal(s.T(), "123", cl.Owner)

	//search
	sl, err := s.clActivity.Search(s.ctx, &pb.ListActivitiesRequest{
		Owner:          "123",
	})
	if err != nil {
		s.T().Fatal(err)
	}
	assert.NotEmpty(s.T(), sl)

	// delete sample
	_, err = s.clActivity.Delete(s.ctx, &pb.ActivityIdRequest{Id: cl.Id})
	if err != nil {
		s.T().Fatal()
	}
}

func (s *activityGrpcTestSuite) TestActivityTypeCRUD() {

	// create a new consultant
	cl, err := s.clActivity.CreateActivityType(s.ctx, s.getCreateActivityTypeRequest())
	if err != nil {
		s.T().Fatal(err)
	}
	assert.NotEmpty(s.T(), cl.Id)

	// get by Id
	cl, err = s.clActivity.GetActivityType(s.ctx, &pb.ActivityTypeIdRequest{Id: cl.Id})
	if err != nil {
		s.T().Fatal()
	}
	assert.NotEmpty(s.T(), cl)
	assert.NotEmpty(s.T(), cl.Id)
	assert.NotEmpty(s.T(), cl.Name)

	// set status to active
	cl, err = s.clActivity.UpdateActivityType(s.ctx, &pb.UpdateActivityTypeRequest{
		Id:          cl.Id,
		Owner: "123",
		Family:     "321",
		Name:     "Monday",
		Description:   "Description",
	})
	if err != nil {
		s.T().Fatal()
	}
	assert.Equal(s.T(), "321", cl.Subject)

	// delete Event
	_, err = s.clActivity.DeleteActivityType(s.ctx, &pb.ActivityTypeIdRequest{Id: cl.Id})
	if err != nil {
		s.T().Fatal(err)
	}
}
