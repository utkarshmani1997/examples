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

func NewAuthBasicUserSignupRequest() *AuthBasicUserSignupRequest {
	return &AuthBasicUserSignupRequest{}
}

func NewAuthBasicUserSignupResponse() *AuthBasicUserSignupResponse {
	return &AuthBasicUserSignupResponse{}
}

func NewAuthBasicUserLoginRequest() *AuthBasicUserLoginRequest {
	return &AuthBasicUserLoginRequest{}
}

func NewAuthBasicUserLoginResponse() *AuthBasicUserLoginResponse {
	return &AuthBasicUserLoginResponse{}
}

// HOFSTADTER_BELOW
