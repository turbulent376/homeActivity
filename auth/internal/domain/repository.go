package domain

// Specify all the contracts your domain logic depends on
// It spans external services and storages

import (
	"context"
)

type UserStorage interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, user *User) error
	// UpdateUser updates an user
	DeleteUser(ctx context.Context, id string) error
	// UpdateUser updates an user
	UpdateUser(ctx context.Context, user *User) error
	// GetUser retrieves a user by id
	GetUser(ctx context.Context, id string) (*User, error)
	// GetUserByEmail gets user by email
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	// GetByFirebaseUUID gets user by firebase auth token
	GetUserByFireUUID(ctx context.Context, uuid string) (*User, error)
}

type SessionStorage interface {
	// CreateUserSession create new user session with deviceId and refresh token
	CreateSession(ctx context.Context, session *Session) error
	// DeleteUserSession delete session(logout)
	DeleteSession(ctx context.Context, id string) error
	// UpdateUserSession update fcm token or refresh token
	UpdateSession(ctx context.Context, session *Session) error
	// GetSession gets session
	GetSession(ctx context.Context, id string) (*Session, error)
	// GetUserSessions get all active user's sessions
	GetUserSessions(ctx context.Context, id string) ([]*Session, error)
}
