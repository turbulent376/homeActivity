package domain

import (
	"context"
)

type ActivityStorage interface {
	// CreateActivity creates a new Activity
	CreateActivity(ctx context.Context, activity *Activity) (*Activity, error)
	// UpdateActivity updates an existent Activity
	UpdateActivity(ctx context.Context, activity *Activity) (*Activity, error)
	// GetActivity retrieves an Activity by id
	GetActivity(ctx context.Context, id string) (bool, *Activity, error)
	// ListActivities lists all Activities owned by user
	ListActivities(ctx context.Context, userId string) (bool, []*Activity, error)
	//ListActivitiesByFamily retrieves all activities done by all Family members
	ListActivitiesByFamily(ctx context.Context, familyId string) (bool, []*Activity, error)
	// Delete deletes an Activity
	DeleteActivity(ctx context.Context, id string) error
}

type ActivityTypeStorage interface {
	// CreateActivityType creates a new ActivityType
	CreateActivityType(ctx context.Context, activityType *ActivityType) (*ActivityType, error)
	// UpdateActivityType updates an existent ActivityType
	UpdateActivityType(ctx context.Context, activityType *ActivityType) (*ActivityType, error)
	// GetActivityType retrieves an ActivityType by id
	GetActivityType(ctx context.Context, id string) (bool, *ActivityType, error)
	// ListActivityTypes lists all Activities owned by Family of the user
	ListActivityTypes(ctx context.Context, familyId string) (bool, []*ActivityType, error)
	// DeleteActivityType deletes an ActivityType by id
	DeleteActivityType(ctx context.Context, id string) error
}
