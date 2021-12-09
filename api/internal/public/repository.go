package public

import (
	authPb "git.jetbrains.space/orbi/fcsd/proto/auth"
	billPb "git.jetbrains.space/orbi/fcsd/proto/billing"
	licPb "git.jetbrains.space/orbi/fcsd/proto/license"
	locPb "git.jetbrains.space/orbi/fcsd/proto/location"
	repPb "git.jetbrains.space/orbi/fcsd/proto/report"
	storePb "git.jetbrains.space/orbi/fcsd/proto/storage"
	timesPb "git.jetbrains.space/orbi/fcsd/proto/timesheet"

	"context"
	"io"
)

type BillingRepository interface {
	MakePayment(ctx context.Context, rq *billPb.MakePaymentRequest) (*billPb.Payment, error)
	GetProducts(ctx context.Context) (*billPb.GetProductsResponse, error)
	CreateProduct(ctx context.Context, rq *billPb.CreateProductRequest) (*billPb.Product, error)
	GetProduct(ctx context.Context, rq *billPb.GetProductRequest) (*billPb.Product, error)
	UpdateProduct(ctx context.Context, rq *billPb.Product) (*billPb.Product, error)
	DeleteProduct(ctx context.Context, rq *billPb.DeleteProductRequest) error
	AppleWebhook(ctx context.Context, rq *billPb.AppleWebhookRequest) error
}

type LicenseRepository interface {
	GetUserLicenses(ctx context.Context, rq *licPb.GetUserLicensesRequest) (*licPb.UserLicenses, error)
	GetLicenses(ctx context.Context) (*licPb.LicenseListResponse, error)
	GetLicense(ctx context.Context, rq *licPb.LicenseRequest) (*licPb.License, error)
	CreateLicense(ctx context.Context, rq *licPb.CreateLicenseRequest) (*licPb.License, error)
	DeleteLicense(ctx context.Context, rq *licPb.LicenseRequest) error
	UpdateLicense(ctx context.Context, rq *licPb.UpdateLicenseRequest) (*licPb.License, error)
}

type ReportRepository interface {
	CreateReport(ctx context.Context, request *repPb.CreateReportRequest) (*repPb.Report, error)
}

type StorageRepository interface {
	PutFile(ctx context.Context, fi *storePb.StorePutFileRequest_FileInfo, file io.Reader) (*storePb.StorePutFileResponse, error)
	GetFile(ctx context.Context, rq *storePb.StoreFileIDRequest) (*storePb.StoreGetFileResponse, error)
	GetMetadata(ctx context.Context, rq *storePb.StoreFileIDRequest) (*storePb.StoreFileInfo, error)
	DeleteFile(ctx context.Context, rq *storePb.StoreFileIDRequest) (*storePb.NoResponse, error)
}

type TimesheetRepository interface {
	CreateTimesheet(ctx context.Context, rq *timesPb.CreateTimesheetRequest) (*timesPb.Timesheet, error)
	UpdateTimesheet(ctx context.Context, rq *timesPb.UpdateTimesheetRequest) (*timesPb.Timesheet, error)
	GetTimesheet(ctx context.Context, rq *timesPb.TimesheetIdRequest) (*timesPb.Timesheet, error)
	SearchTimesheet(ctx context.Context, rq *timesPb.SearchTimesheetRequest) (*timesPb.Timesheet, error)
	DeleteTimesheet(ctx context.Context, rq *timesPb.TimesheetIdRequest) error
	CreateEvent(ctx context.Context, rq *timesPb.CreateEventRequest) (*timesPb.Event, error)
	UpdateEvent(ctx context.Context, rq *timesPb.UpdateEventRequest) (*timesPb.Event, error)
	GetEvent(ctx context.Context, rq *timesPb.EventIdRequest) (*timesPb.Event, error)
	DeleteEvent(ctx context.Context, rq *timesPb.EventIdRequest) error
	SearchEvents(ctx context.Context, rq *timesPb.EventIdRequest) (*timesPb.SearchResponse, error)
}

type LocationRepository interface {
	GetAllLocations(ctx context.Context, request *locPb.PagingRequest) (*locPb.LocationsResponse, error)
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
