package auth

import pb "git.jetbrains.space/orbi/fcsd/proto/auth"

func (c *ctrlImpl) toAuthRequestPb(rq *AuthRequest) *pb.AuthRequest {
	return &pb.AuthRequest{
		Email:    rq.Email,
		Password: rq.Password,
	}
}

func (c *ctrlImpl) toAuthResponse(rq *pb.AuthResponse) *AuthResponse {
	return &AuthResponse{
		Token:        rq.Token,
		RefreshToken: rq.RefreshToken,
		User:         c.toUserResponse(rq.User),
	}
}

func (c *ctrlImpl) toTokenPairResponse(rq *pb.TokenPairResponse) *TokenPairResponse {
	return &TokenPairResponse{
		Token:        rq.Token,
		RefreshToken: rq.RefreshToken,
	}
}

func (c *ctrlImpl) toCreateUserPb(rq *CreateUserRequest) *pb.CreateUserRequest {
	return &pb.CreateUserRequest{
		Name:     rq.Name,
		Surname:  rq.Surname,
		Email:    rq.Email,
		Password: rq.Password,
	}
}

func (c *ctrlImpl) toUserInfoResponse(rq *pb.User) *UserInfoResponse {
	return &UserInfoResponse{
		Id:      rq.Id,
		Name:    rq.Name,
		Surname: rq.Surname,
		Avatar:  rq.Avatar,
	}
}

func (c *ctrlImpl) toUserResponse(rq *pb.User) *UserResponse {
	return &UserResponse{
		Id:          rq.Id,
		Email:       rq.Email,
		Name:        rq.Name,
		Surname:     rq.Surname,
		Avatar:      rq.Avatar,
		CountryCode: rq.CountryCode,
	}
}

func (c *ctrlImpl) toUpdateUserRequestPb(rq *UpdateUserRequest) *pb.UpdateUserRequest {
	return &pb.UpdateUserRequest{
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
