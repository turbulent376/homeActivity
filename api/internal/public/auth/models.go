package auth

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OAuthRequest struct {
	Token string `json:"token"`
}

type TokenPairResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type AuthResponse struct {
	Token        string        `json:"token"`
	RefreshToken string        `json:"refreshToken"`
	User         *UserResponse `json:"user"`
}
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type LogoutRequest struct {
	SessionId string `json:"sessionId"`
}

type FCMTokenRequest struct {
	FCMToken string `json:"fcmToken"`
}

type UserResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	CountryCode string `json:"countryCode"`
}

// user's public info
type UserInfoResponse struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Avatar  string `json:"avatar"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
	CountryCode string `json:"countryCode"`
}
