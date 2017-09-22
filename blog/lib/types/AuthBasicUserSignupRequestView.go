package types

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      AuthBasicUserSignupRequest
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type AuthBasicUserSignupRequest struct {
	Password string `json:"password" xml:"password" yaml:"password" form:"password" query:"password" validate:"required" `

	PasswordVerify string `json:"password_verify" xml:"password_verify" yaml:"password_verify" form:"password_verify" query:"password_verify" validate:"required" `

	Email string `json:"email" xml:"email" yaml:"email" form:"email" query:"email" validate:"required|email|min=1|max=64" `

	Username string `json:"username" xml:"username" yaml:"username" form:"username" query:"username" validate:"required|alphanumunicode|min=1|max=64" `
}

// HOFSTADTER_BELOW
