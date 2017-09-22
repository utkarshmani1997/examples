package types

import (
	"github.com/jinzhu/gorm"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
Name:      AuthBasicUser
About:     Basic Auth type
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type AuthBasicUser struct {

	/* ORM: server.api.databases.[name==postgres]*/

	gorm.Model
	UUID string `json:"uuid" xml:"uuid" yaml:"uuid" form:"uuid" query:"uuid" validate:"uuidv4" `

	password string `json:"password" xml:"password" yaml:"password" form:"password" query:"password" validate:"required" `

	// Relations ----
	// Should be using some type lookup and package resolution type things

	// owned-by (the other half of has-one/many)
	UserID uint `json:"user_id" xml:"user_id" yaml:"user_id" form:"user_id" query:"user_id" `
}

// HOFSTADTER_BELOW
