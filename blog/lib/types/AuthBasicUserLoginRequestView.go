package types

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      AuthBasicUserLoginRequest
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type AuthBasicUserLoginRequest struct {
	Password string `json:"password" xml:"password" yaml:"password" form:"password" query:"password" validate:"required" `

	Email string `json:"email" xml:"email" yaml:"email" form:"email" query:"email" validate:"required|email|min=1|max=64" `
}

// HOFSTADTER_BELOW
