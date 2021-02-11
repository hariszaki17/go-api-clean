package model

// CreateUserRequest expose global
type CreateUserRequest struct {
	Username	string
	Password	string
	Role		string
}

// CreateUserResponse expose global
type CreateUserResponse struct {
	Username	string
	Role		string
}

// GetUserResponse expose global
type GetUserResponse struct {
	ID			interface{}
	Username	string
	Role		string
}

// DeleteAllUserResponse expose global
type DeleteAllUserResponse struct {
	Message		string
}

// LoginUserRequest expose global
type LoginUserRequest struct {
	Username	string
	Password	string
}

// LoginUserResponse expose global
type LoginUserResponse struct {
	Message	string
}