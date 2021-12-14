package activity

import (
	pb "github.com/turbulent376/proto/activity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *ctrlImpl) toCreateActivityPb(rq *CreateActivityRequest) *pb.CreateActivityRequest {
	return &pb.CreateActivityRequest{
		Owner:    rq.Owner,
		Family:   rq.Family,
		Type:     rq.Type,
		DateFrom: timestamppb.New(rq.DateFrom),
		DateTo:   timestamppb.New(rq.DateTo),
	}
}

func (c *ctrlImpl) toActivityApi(rq *pb.Activity) *Activity {
	return &Activity{
		Id:       rq.Id,
		Owner:    rq.Owner,
		Family:   rq.Family,
		Type:     rq.Type,
		DateFrom: rq.DateFrom.AsTime(),
		DateTo:   rq.DateTo.AsTime(),
	}
}

func (c *ctrlImpl) toActivitiesApi(rq *pb.ListActivitiesResponse) ListActivitiesResponse {
	var res []*Activity
	for _, result := range rq.GetActivities() {
		res = append(res, c.toActivityApi(result))
	}
	return ListActivitiesResponse{Result: res}
}

func (c *ctrlImpl) toUpdateActivityPb(rq *UpdateActivityRequest) *pb.UpdateActivityRequest {
	return &pb.UpdateActivityRequest{
		Id:      rq.Id,
		Owner:    rq.Owner,
		Family:   rq.Family,
		Type:     rq.Type,
		DateFrom: timestamppb.New(rq.DateFrom),
		DateTo: timestamppb.New(rq.DateTo),
	}
}

func (c *ctrlImpl) toCreateActivityTypePb(rq *CreateActivityTypeRequest) *pb.CreateActivityTypeRequest {
	return &pb.CreateActivityTypeRequest{
		Family:       rq.Family,
		Name:         rq.Name,
		Description:  rq.Description,
	}
}

func (c *ctrlImpl) toActivityTypeApi(rq *pb.ActivityType) *ActivityType {
	return &ActivityType{
		Id:           rq.Id,
		Family:       rq.Family,
		Name:         rq.Name,
		Description:  rq.Description,
	}
}

func (c *ctrlImpl) toUpdateActivityTypePb(rq *UpdateActivityTypeRequest) *pb.UpdateActivityTypeRequest {
	return &pb.UpdateActivityTypeRequest{
		Id:          rq.Id,
		Family:       rq.Family,
		Name:         rq.Name,
		Description:  rq.Description,
	}
}

func (c *ctrlImpl) toActivityTypesApi(rq *pb.ListActivityTypesResponse) ListActivityTypesResponse {
	var res []*ActivityType
	for _, result := range rq.GetActyvityTypes() {
		res = append(res, c.toActivityTypeApi(result))
	}
	return ListActivityTypesResponse{Result: res}
}
