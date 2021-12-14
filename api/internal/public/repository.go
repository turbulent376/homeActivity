package public

import (
	authPb "github.com/turbulent376/proto/auth"
	activPb "github.com/turbulent376/proto/activity"

	"context"
	//"io"
)

type ActivityRepository interface {
	CreateActivity(ctx context.Context, rq *activPb.CreateActivityRequest) (*activPb.Activity, error)
	UpdateActivity(ctx context.Context, rq *activPb.UpdateActivityRequest) (*activPb.Activity, error)
	GetActivity(ctx context.Context, rq *activPb.ActivityIdRequest) (*activPb.Activity, error)
	ListActivities(ctx context.Context, rq *activPb.ListActivitiesRequest) (*activPb.ListActivitiesResponse, error)
	ListActivitiesByFamily(ctx context.Context, rq *activPb.ListActivitiesByFamilyRequest) (*activPb.ListActivitiesResponse, error)
	DeleteActivity(ctx context.Context, rq *activPb.ActivityIdRequest) error
	CreateActivityType(ctx context.Context, rq *activPb.CreateActivityTypeRequest) (*activPb.ActivityType, error)
	UpdateActivityType(ctx context.Context, rq *activPb.UpdateActivityTypeRequest) (*activPb.ActivityType, error)
	GetActivityType(ctx context.Context, rq *activPb.ActivityTypeIdRequest) (*activPb.ActivityType, error)
	DeleteActivityType(ctx context.Context, rq *activPb.ActivityTypeIdRequest) error
	ListActivityTypes(ctx context.Context, rq *activPb.ListActivityTypesRequest) (*activPb.ListActivityTypesResponse, error)
}

type AuthRepository interface {
	AuthUserByEmail(ctx context.Context, rq *authPb.AuthRequest) (*authPb.AuthResponse, error)
	AuthUserByFirebase(ctx context.Context, rq *authPb.OAuthRequest) (*authPb.AuthResponse, error)
	RefreshToken(ctx context.Context, rq *authPb.RefreshTokenRequest) (*authPb.TokenPairResponse, error)
	GetUserById(ctx context.Context, rq *authPb.UserIdRequest) (*authPb.User, error)
	GetSessionByToken(ctx context.Context, rq *authPb.TokenRequest) (*authPb.Session, error)
	CreateUser(ctx context.Context, rq *authPb.CreateUserRequest) (*authPb.User, error)
	DeleteUser(ctx context.Context, rq *authPb.UserIdRequest) error
	UpdateUser(ctx context.Context, rq *authPb.UpdateUserRequest) (*authPb.User, error)
	CloseSession(ctx context.Context, rq *authPb.CloseSessionRequest) error
	SaveUserFCMToken(ctx context.Context, rq *authPb.FCMTokenRequest) error
}
