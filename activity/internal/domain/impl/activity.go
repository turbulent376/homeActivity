package impl

import (
	"context"
	"github.com/turbulent376/homeactivity/activity/internal/domain"
	"github.com/turbulent376/homeactivity/activity/internal/errors"
	"github.com/turbulent376/homeactivity/activity/internal/logger"
	"github.com/turbulent376/kit/log"
	"github.com/turbulent376/kit/utils"
)

// sampleTimesheetImpl - implements SampleService interface
type activitytImpl struct {
	storageA domain.ActivityStorage
	storageT domain.ActivityTypeStorage
}

func NewActivityService(
	storageA domain.ActivityStorage,
	storageT domain.ActivityTypeStorage,
) domain.ActivityService {
	return &activitytImpl{
		storageA: storageA,
		storageT: storageT,
	}
}

func (s *activitytImpl) l() log.CLogger {
	return logger.L().Cmp("timesheet")
}

func (s *activitytImpl) Create(ctx context.Context, ac *domain.Activity) (*domain.Activity, error) {
	s.l().C(ctx).Mth("create").Dbg()

	// check owner
	if ac.Owner == "" {
		return nil, errors.ErrActivityOwnerIsEmpty(ctx)
	}

	//check type
	if ac.Type == "" {
		return nil, errors.ErrActivityTypeIsEmpty(ctx)
	}

	//check family
	if ac.Family == ""{
		return nil, errors.ErrActivityFamilyIdIsEmpty(ctx)
	}

	// check timeFrom
	if ac.DateFrom.IsZero() {
		return nil, errors.ErrActivityTimeIsEmpty(ctx)
	}

	// check timeTo
	if ac.DateTo.IsZero() {
		return nil, errors.ErrActivityTimeIsEmpty(ctx)
	}

	ac.Id = utils.NewId()

	// save to store
	ac, err := s.storageA.CreateActivity(ctx, ac)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

func (s *activitytImpl) Update(ctx context.Context, ac *domain.Activity) (*domain.Activity, error) {
	s.l().C(ctx).Mth("update").Dbg()

	// validates id
	if ac.Id == "" {
		return nil, errors.ErrActivityIdIsEmpty(ctx)
	}

	// check owner
	if ac.Owner == "" {
		return nil, errors.ErrActivityOwnerIsEmpty(ctx)
	}

	//check type
	if ac.Type == "" {
		return nil, errors.ErrActivityTypeIsEmpty(ctx)
	}

	//check family
	if ac.Family == ""{
		return nil, errors.ErrActivityFamilyIdIsEmpty(ctx)
	}

	// check timeFrom
	if ac.DateFrom.IsZero() {
		return nil, errors.ErrActivityTimeIsEmpty(ctx)
	}

	// check timeTo
	if ac.DateTo.IsZero() {
		return nil, errors.ErrActivityTimeIsEmpty(ctx)
	}


	// retrieve stored sample by id
	found, stored, err := s.storageA.GetActivity(ctx, ac.Id)

	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.ErrActivityNotFound(ctx,ac.Id)
	}

	// set updated params
	stored.Type = ac.Type
	stored.DateFrom = ac.DateFrom
	stored.DateTo = ac.DateTo

	// save to store
	_, err = s.storageA.UpdateActivity(ctx, stored)
	if err != nil {
		return nil, err
	}

	return stored, nil
}

func (s *activitytImpl) Get(ctx context.Context, id string) (bool, *domain.Activity, error) {
	s.l().C(ctx).Mth("get").Dbg()
	// validates id
	if id == "" {
		return false, nil, errors.ErrActivityIdIsEmpty(ctx)
	}
	return s.storageA.GetActivity(ctx, id)
}

func (s *activitytImpl) ListActivities(ctx context.Context, userId string) (bool, []*domain.Activity, error) {
	s.l().C(ctx).Mth("ListActivities").Dbg()
	if userId == "" {
		return false, nil, errors.ErrActivityUserIdIsEmpty(ctx)
	}
	return s.storageA.ListActivities(ctx, userId)
}

func (s *activitytImpl) ListActivitiesByFamily(ctx context.Context, familyId string) (bool, []*domain.Activity, error){
	s.l().C(ctx).Mth("ListActivitiesByFamily").Dbg()

	if familyId == "" {
		return false, nil, errors.ErrActivityFamilyIdIsEmpty(ctx)
	}
	return s.storageA.ListActivitiesByFamily(ctx, familyId)
}

func (s *activitytImpl) Delete(ctx context.Context, id string) error {
	s.l().C(ctx).Mth("delete").Dbg()

	// check id isn't empty
	if id == "" {
		return errors.ErrActivityOwnerIsEmpty(ctx)
	}

	// retrieve stored sample by id
	found, _, err := s.storageA.GetActivity(ctx, id)
	if err != nil {
		return err
	}
	if !found {
		return errors.ErrActivityNotFound(ctx, id)
	}

	// delete from storage
	return s.storageA.DeleteActivity(ctx, id)
}

func (s *activitytImpl) CreateActivityType(ctx context.Context, at *domain.ActivityType) (*domain.ActivityType, error) {
	s.l().C(ctx).Mth("create").Dbg()

	if at.Family == "" {
		return nil, errors.ErrActivityFamilyIdIsEmpty(ctx)
	}
	if at.Name == "" {
		return nil, errors.ErrActivityNameIsEmpty(ctx)
	}
	if at.Description == "" {
		return nil, errors.ErrActivityDescriptionIsEmpty(ctx)
	}

	at.Id = utils.NewId()

	// save to store
	ev, err := s.storageT.CreateActivityType(ctx, at)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

func (s *activitytImpl) UpdateActivityType(ctx context.Context, at *domain.ActivityType) (*domain.ActivityType, error) {
	s.l().C(ctx).Mth("UpdateActivityType").Dbg()

	// check timesheet isn't empty
	if at.Id == "" {
		return nil, errors.ErrActivityIdIsEmpty(ctx)
	}
	if at.Family == "" {
		return nil, errors.ErrActivityFamilyIdIsEmpty(ctx)
	}
	if at.Name == "" {
		return nil, errors.ErrActivityNameIsEmpty(ctx)
	}
	if at.Description == "" {
		return nil, errors.ErrActivityDescriptionIsEmpty(ctx)
	}

	// retrieve stored sample by id
	found, stored, err := s.storageT.GetActivityType(ctx, at.Id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.ErrActivityNotFound(ctx, at.Id)
	}

	// set updated params
	stored.Family = at.Family
	stored.Name = at.Name
	stored.Description = at.Description

	// save to store
	_, err = s.storageT.UpdateActivityType(ctx, stored)
	if err != nil {
		return nil, err
	}

	return stored, nil
}

func (s *activitytImpl) GetActivityType(ctx context.Context, id string) (bool, *domain.ActivityType, error) {
	s.l().C(ctx).Mth("GetActivityType").Dbg()
	if id == "" {
		return false, nil, errors.ErrActivityIdIsEmpty(ctx)
	}
	return s.storageT.GetActivityType(ctx, id)
}

func (s *activitytImpl) ListActivityTypes(ctx context.Context, familyId string) (bool, []*domain.ActivityType, error){
	s.l().C(ctx).Mth("ListActivityTypes").Dbg()
	if familyId == "" {
		return false, nil, errors.ErrActivityFamilyIdIsEmpty(ctx)
	}
	return s.storageT.ListActivityTypes(ctx, familyId)
}

func (s *activitytImpl) DeleteActivityType(ctx context.Context, id string) error {
	s.l().C(ctx).Mth("delete").Dbg()

	// check id isn't empty
	if id == "" {
		return errors.ErrActivityIdIsEmpty(ctx)
	}

	// retrieve stored sample by id
	found, _, err := s.storageT.GetActivityType(ctx, id)
	if err != nil {
		return err
	}
	if !found {
		return errors.ErrActivityNotFound(ctx, id)
	}

	// save to store
	return s.storageT.DeleteActivityType(ctx, id)
}
