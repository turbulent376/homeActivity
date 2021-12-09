package grpc

import (
	"context"

	pb "git.jetbrains.space/orbi/fcsd/proto/auth"
)

func (s *Server) CreateUser(ctx context.Context, rq *pb.CreateUserRequest) (*pb.User, error) {
	sample, err := s.authService.CreateUser(ctx, s.toCreateUserDomain(rq))
	if err != nil {
		return nil, err
	}
	return s.toUserPb(sample), nil
}

func (s *Server) UpdateUser(ctx context.Context, rq *pb.UpdateUserRequest) (*pb.User, error) {
	user, err := s.authService.UpdateUser(ctx, s.toUpdateUserDomain(rq))
	if err != nil {
		return nil, err
	}
	return s.toUserPb(user), nil
}

func (s *Server) GetUserById(ctx context.Context, rq *pb.UserIdRequest) (*pb.User, error) {
	user, err := s.authService.GetUserById(ctx, rq.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return s.toUserPb(user), nil
}

func (s *Server) GetSessionByToken(ctx context.Context, rq *pb.TokenRequest) (*pb.Session, error) {
	session, err := s.authService.GetSessionByToken(ctx, rq.Token)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, nil
	}
	return s.toSessionPb(session), nil
}

func (s *Server) DeleteUser(ctx context.Context, rq *pb.UserIdRequest) (*pb.EmptyResponse, error) {
	err := s.authService.DeleteUser(ctx, rq.Id)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResponse{}, nil
}

func (s *Server) AuthUserByEmail(ctx context.Context, rq *pb.AuthRequest) (*pb.AuthResponse, error) {
	resp, err := s.authService.AuthUserByEmail(ctx, s.toAuthRequestDomain(rq))
	if err != nil {
		return nil, err
	}
	return s.toAuthResponsePb(resp), nil
}

func (s *Server) AuthUserByFirebase(ctx context.Context, rq *pb.OAuthRequest) (*pb.AuthResponse, error) {
	resp, err := s.authService.AuthUserByFirebase(ctx, s.toOAuthRequestDomain(rq))
	if err != nil {
		return nil, err
	}
	return s.toAuthResponsePb(resp), nil
}

func (s *Server) RefreshToken(ctx context.Context, rq *pb.RefreshTokenRequest) (*pb.TokenPairResponse, error) {
	tokPair, err := s.authService.RefreshToken(ctx, s.toRefreshTokenRequestDomain(rq))
	if err != nil {
		return nil, err
	}
	return s.toTokenPairResponsePb(tokPair), nil
}

func (s *Server) SaveUserFCMToken(ctx context.Context, rq *pb.FCMTokenRequest) (*pb.EmptyResponse, error) {
	err := s.authService.SaveUserFCMToken(ctx, s.toFCMTokenRequestDomain(rq))
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResponse{}, nil
}

func (s *Server) CloseSession(ctx context.Context, rq *pb.CloseSessionRequest) (*pb.EmptyResponse, error) {
	err := s.authService.CloseSession(ctx, s.toCloseSessionRequestDomain(rq))
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResponse{}, nil
}

func (s *Server) GetUserFCMTokens(ctx context.Context, rq *pb.UserIdRequest) (*pb.FCMTokensResponse, error) {
	tokens, err := s.authService.GetUserFCMTokens(ctx, rq.Id)
	if err != nil {
		return nil, err
	}
	return &pb.FCMTokensResponse{FCMTokens: tokens}, nil
}

func (s *Server) GetUserSessions(ctx context.Context, rq *pb.UserIdRequest) (*pb.UserSessionsResponse, error) {
	sessions, err := s.authService.GetUserSessions(ctx, rq.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UserSessionsResponse{Sessions: s.toSessionArrayPb(sessions)}, nil
}
