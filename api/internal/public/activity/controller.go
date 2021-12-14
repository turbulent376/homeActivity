package activity

import (
	"github.com/turbulent376/homeactivity/api/internal/public"
	kitHttp "github.com/turbulent376/kit/http"
	pb "github.com/turbulent376/proto/activity"
	"net/http"
)

type Controller interface {
	CreateActivity(w http.ResponseWriter, r *http.Request)
	UpdateActivity(w http.ResponseWriter, r *http.Request)
	GetActivity(w http.ResponseWriter, r *http.Request)
	ListActivities(w http.ResponseWriter, r *http.Request)
	ListActivitiesByFamily(w http.ResponseWriter, r *http.Request)
	DeleteActivity(w http.ResponseWriter, r *http.Request)
	CreateActivityType(w http.ResponseWriter, r *http.Request)
	UpdateActivityType(w http.ResponseWriter, r *http.Request)
	GetActivityType(w http.ResponseWriter, r *http.Request)
	DeleteActivityType(w http.ResponseWriter, r *http.Request)
	ListActivityTypes(w http.ResponseWriter, r *http.Request)
}

type ctrlImpl struct {
	kitHttp.BaseController
	activityRepo public.ActivityRepository
}

func NewController(ar public.ActivityRepository) Controller {
	return &ctrlImpl{
		activityRepo: ar,
	}
}

// CreateActivity godoc
// @Summary create activity
// @Accept json
// @Router /activity/activity [POST]
// @Param json body CreateTimesheetRequest true "request"
// @Produce json
// @Success 200 {object} Timesheet
// @Failure 500 {object} kitHttp.Error
// @tags timesheet
func (c *ctrlImpl) CreateActivity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request *CreateActivityRequest

	err := c.DecodeRequest(r, ctx, &request)

	if err != nil {
		c.RespondError(w, err)
		return
	}
	activity, err := c.activityRepo.CreateActivity(ctx, c.toCreateActivityPb(request))
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toActivityApi(activity))
}

// UpdateActivity godoc
// @Summary update activity
// @Accept json
// @Router /activity/activity/{id} [PUT]
// @Param json body UpdateActivityRequest true "request"
// @Param id path string true "id of activity"
// @Produce json
// @Success 200 {object} Activity
// @Failure 500 {object} kitHttp.Error
// @tags activity
func (c *ctrlImpl) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request *UpdateActivityRequest

	id, err := c.Var(r, ctx, "id", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	err = c.DecodeRequest(r, ctx, &request)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	request.Id = id

	activity, err := c.activityRepo.UpdateActivity(ctx, c.toUpdateActivityPb(request))
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toActivityApi(activity))
}

// GetActivity godoc
// @Summary getting activity
// @Router /activity/activity/{id} [GET]
// @Param id path string true "id"
// @Produce json
// @Success 200 {object} Activity
// @Failure 500 {object} kitHttp.Error
// @tags activity
func (c *ctrlImpl) GetActivity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := c.Var(r, ctx, "id", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	activity, err := c.activityRepo.GetActivity(ctx, &pb.ActivityIdRequest{Id: id})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toActivityApi(activity))
}

// ListActivities godoc
// @Summary getting activities by owner
// @Router /activity/list/{owner} [GET]
// @Param owner path string true "owner"
// @Produce json
// @Success 200 {object} ListActivitiesResponse
// @Failure 500 {object} kitHttp.Error
// @tags activity
func (c *ctrlImpl) ListActivities(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	owner, err := c.Var(r, ctx, "owner", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	resp, err := c.activityRepo.ListActivities(ctx, &pb.ListActivitiesRequest{
		Owner:          owner,
	})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toActivitiesApi(resp))
}

// ListActivitiesByFamily godoc
// @Summary getting activities by family
// @Router /activity/listfamily/{family} [GET]
// @Param owner path string true "family"
// @Produce json
// @Success 200 {object} ListActivitiesResponse
// @Failure 500 {object} kitHttp.Error
// @tags activity
func (c *ctrlImpl) ListActivitiesByFamily(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	family, err := c.Var(r, ctx, "family", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	resp, err := c.activityRepo.ListActivitiesByFamily(ctx, &pb.ListActivitiesByFamilyRequest{
		Family:          family,
	})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toActivitiesApi(resp))
}

// DeleteActivity godoc
// @Summary delete activity
// @Router /activity/activity/{id} [DELETE]
// @Param id path string true "id of activity"
// @Success 200
// @Failure 500 {object} kitHttp.Error
// @tags activity
func (c *ctrlImpl) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := c.Var(r, ctx, "id", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	err = c.activityRepo.DeleteActivity(ctx, &pb.ActivityIdRequest{Id: id})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, nil)
}

// CreateActivityType godoc
// @Summary create activityType
// @Accept json
// @Router /activity/activitytype [POST]
// @Param json body CreateActivityTypeRequest true "request"
// @Produce json
// @Success 200 {object} ActivityType
// @Failure 500 {object} kitHttp.Error
// @tags activitytype
func (c *ctrlImpl) CreateActivityType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request *CreateActivityTypeRequest

	err := c.DecodeRequest(r, ctx, &request)

	if err != nil {
		c.RespondError(w, err)
		return
	}
	at, err := c.activityRepo.CreateActivityType(ctx, c.toCreateActivityTypePb(request))
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toActivityTypeApi(at))
}

// UpdateActivityType godoc
// @Summary update activityType
// @Accept json
// @Router /activity/activitytype/{id} [PUT]
// @Param json body UpdateActivityTypeRequest true "request"
// @Param id path string true "id of activityType"
// @Produce json
// @Success 200 {object} ActivityType
// @Failure 500 {object} kitHttp.Error
// @tags activitytype
func (c *ctrlImpl) UpdateActivityType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request *UpdateActivityTypeRequest

	id, err := c.Var(r, ctx, "id", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	err = c.DecodeRequest(r, ctx, &request)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	request.Id = id
	at, err := c.activityRepo.UpdateActivityType(ctx, c.toUpdateActivityTypePb(request))
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toActivityTypeApi(at))
}

// GetActivityType godoc
// @Summary getting activityType
// @Router /activity/activitytype/{id} [GET]
// @Param id path string true "id of activityType"
// @Produce json
// @Success 200 {object} ActivityType
// @Failure 500 {object} kitHttp.Error
// @tags activitytype
func (c *ctrlImpl) GetActivityType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := c.Var(r, ctx, "id", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	at, err := c.activityRepo.GetActivityType(ctx, &pb.ActivityTypeIdRequest{
		Id: id,
	})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toActivityTypeApi(at))
}

// DeleteActivityType godoc
// @Summary delete activityType
// @Router /activity/activitytype/{id} [DELETE]
// @Param id path string true "id of activityType"
// @Success 200
// @Failure 500 {object} kitHttp.Error
// @tags activitytype
func (c *ctrlImpl) DeleteActivityType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := c.Var(r, ctx, "id", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	err = c.activityRepo.DeleteActivityType(ctx, &pb.ActivityTypeIdRequest{Id: id})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, nil)
}

// ListActivityTypes godoc
// @Summary list activityTypes by family
// @Router /activity/listactivitytypes/{family} [GET]
// @Param family path string true "id of family"
// @Produce json
// @Success 200 {object} ListActivityTypesResponse
// @Failure 500 {object} kitHttp.Error
// @tags activitytype
func (c *ctrlImpl) ListActivityTypes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	family, err := c.Var(r, ctx, "family", false)

	if err != nil {
		c.RespondError(w, err)
		return
	}

	resp, err := c.activityRepo.ListActivityTypes(ctx, &pb.ListActivityTypesRequest{Family: family})
	if err != nil {
		c.RespondError(w, err)
		return
	}
	c.RespondOK(w, c.toActivityTypesApi(resp))
}
