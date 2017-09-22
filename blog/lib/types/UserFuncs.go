package types

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      User
About:     A user of the blog site
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

func NewUser() *User {
	return &User{}
}

func NewUserViewAuthBasicUserSignupRequest() *UserViewAuthBasicUserSignupRequest {
	return &UserViewAuthBasicUserSignupRequest{}
}

func NewUserViewAuthBasicUserSignupResponse() *UserViewAuthBasicUserSignupResponse {
	return &UserViewAuthBasicUserSignupResponse{}
}

func NewUserViewAuthBasicUserLoginRequest() *UserViewAuthBasicUserLoginRequest {
	return &UserViewAuthBasicUserLoginRequest{}
}

func NewUserViewAuthBasicUserLoginResponse() *UserViewAuthBasicUserLoginResponse {
	return &UserViewAuthBasicUserLoginResponse{}
}

// HOFSTADTER_BELOW
