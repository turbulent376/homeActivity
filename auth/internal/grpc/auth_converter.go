package grpc

import (
	"git.jetbrains.space/orbi/fcsd/auth/internal/domain"
	pb "git.jetbrains.space/orbi/fcsd/proto/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) toCreateUserDomain(rq *pb.CreateUserRequest) *domain.RegistrationUserRequest {
	return &domain.RegistrationUserRequest{
		Name:     rq.Name,
		Surname:  rq.Surname,
		Email:    rq.Email,
		Password: rq.Password,
	}
}

func (s *Server) toUpdateUserDomain(rq *pb.UpdateUserRequest) *domain.UpdateUserRequest {
	return &domain.UpdateUserRequest{
		Id:          rq.Id,
		Name:        rq.Name,
		Surname:     rq.Surname,
		Email:       rq.Email,
		OldPassword: rq.OldPassword,
		NewPassword: rq.NewPassword,
		Avatar:      rq.Avatar,
		CountryCode: rq.CountryCode,
	}
}

func (s *Server) toUserPb(rq *domain.User) *pb.User {
	return &pb.User{
		Id:          rq.Id,
		Name:        rq.Name,
		Surname:     rq.Surname,
		Avatar:      rq.Avatar,
		Email:       rq.Email,
		CountryCode: rq.CountryCode,
	}
}

func (s *Server) toSessionPb(rq *domain.Session) *pb.Session {
	return &pb.Session{
		Id:            rq.Id,
		UserId:        rq.UserId,
		DeviceName:    rq.DeviceName,
		ClientVersion: rq.ClientVersion,
		RefreshToken:  rq.RefreshToken,
		FCMToken:      rq.FCMToken,
		CreatedAt:     timestamppb.New(rq.CreatedAt),
	}
}

func (s *Server) toAuthRequestDomain(rq *pb.AuthRequest) *domain.AuthRequest {
	return &domain.AuthRequest{
		Email:    rq.Email,
		Password: rq.Password,
	}
}

func (s *Server) toOAuthRequestDomain(rq *pb.OAuthRequest) *domain.OAuthRequest {
	return &domain.OAuthRequest{
		Token: rq.Token,
	}
}

func (s *Server) toRefreshTokenRequestDomain(rq *pb.RefreshTokenRequest) *domain.RefreshTokenRequest {
	return &domain.RefreshTokenRequest{
		UserId:       rq.UserId,
		RefreshToken: rq.RefreshToken,
		SessionId:    rq.SessionId,
	}
}

func (s *Server) toTokenPairResponsePb(rq *domain.TokenPair) *pb.TokenPairResponse {
	return &pb.TokenPairResponse{
		Token:        rq.Token,
		RefreshToken: rq.RefreshToken,
	}
}

func (s *Server) toAuthResponsePb(rq *domain.AuthResponse) *pb.AuthResponse {
	return &pb.AuthResponse{
		Token:        rq.Token,
		RefreshToken: rq.RefreshToken,
		User:         s.toUserPb(&rq.User),
	}
}

func (s *Server) toFCMTokenRequestDomain(rq *pb.FCMTokenRequest) *domain.FCMTokenRequest {
	return &domain.FCMTokenRequest{
		UserId:    rq.UserId,
		SessionId: rq.SessionId,
		FCMToken:  rq.FCMToken,
	}
}

func (s *Server) toCloseSessionRequestDomain(rq *pb.CloseSessionRequest) *domain.CloseSessionRequest {
	return &domain.CloseSessionRequest{
		UserId:    rq.UserId,
		SessionId: rq.SessionId,
	}
}

func (s *Server) toSessionArrayPb(arr []*domain.Session) []*pb.Session {
	var result []*pb.Session

	for _, value := range arr {
		result = append(result, &pb.Session{
			Id:            value.Id,
			UserId:        value.UserId,
			DeviceName:    value.DeviceName,
			RefreshToken:  value.RefreshToken,
			FCMToken:      value.FCMToken,
			ClientVersion: value.ClientVersion,
			CreatedAt:     timestamppb.New(value.CreatedAt),
		})
	}

	return result
}
