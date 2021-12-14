package storage

import (
	"github.com/turbulent376/homeactivity/activity/internal/domain"
)

func (a *adapterImpl) toActivityDto(t *domain.Activity) *activity {
	if t == nil {
		return nil
	}

	dto := &activity{
		Id:       t.Id,
		Owner:    t.Owner,
		Type:     t.Type,
		Family:   t.Family,
		DateFrom: t.DateFrom,
		DateTo:   t.DateTo,
	}

	return dto
}

func (a *adapterImpl) toActivityDomain(dto *activity) *domain.Activity {
	if dto == nil {
		return nil
	}
	return &domain.Activity{
		Id:        dto.Id,
		Owner:     dto.Owner,
		Type:      dto.Type,
		Family:    dto.Family,
		DateFrom:  dto.DateFrom,
		DateTo:    dto.DateTo,
	}
}

func (a *adapterImpl) toActivitiesDomain(dtos []*activity) []*domain.Activity {
	var r []*domain.Activity
	for _, dto := range dtos {
		r = append(r, a.toActivityDomain(dto))
	}
	return r
}

func (a *adapterImpl) toActivityTypeDto(t *domain.ActivityType) *activityType {
	if t == nil {
		return nil
	}

	dto := &activityType{
		Id:          t.Id,
		Family: 	 t.Family,
		Name:        t.Name,
		Description: t.Description,
	}

	return dto
}

func (a *adapterImpl) toActivityTypeDomain(dto *activityType) *domain.ActivityType {
	if dto == nil {
		return nil
	}
	return &domain.ActivityType{
		Id:          dto.Id,
		Family:      dto.Family,
		Name:        dto.Name,
		Description: dto.Description,
	}
}

func (a *adapterImpl) toActivityTypesDomain(dtos []*activityType) []*domain.ActivityType {
	var r []*domain.ActivityType
	for _, dto := range dtos {
		r = append(r, a.toActivityTypeDomain(dto))
	}
	return r
}
