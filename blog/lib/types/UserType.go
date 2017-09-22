package types

import (
	"github.com/jinzhu/gorm"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
Name:      User
About:     A user of the blog site
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type User struct {

	/* ORM: server.api.databases.[name==postgres]*/

	gorm.Model
	UUID string `json:"uuid" xml:"uuid" yaml:"uuid" form:"uuid" query:"uuid" validate:"uuidv4" `

	Email string `json:"email" xml:"email" yaml:"email" form:"email" query:"email" validate:"required|email|min=1|max=64" `

	Username string `json:"username" xml:"username" yaml:"username" form:"username" query:"username" validate:"required|alphanumunicode|min=1|max=64" `

	// Relations ----
	// Should be using some type lookup and package resolution type things

	// has-many
	Posts []Post `json:"posts" xml:"posts" yaml:"posts" form:"posts" query:"posts" `

	// many-to-many
	Likes []Post `json:"likes" xml:"likes" yaml:"likes" form:"likes" query:"likes" gorm:"many2many:user_likes"`

	// has-one
	AuthBasicUser AuthBasicUser `json:"auth_basic_user" xml:"auth_basic_user" yaml:"auth_basic_user" form:"auth_basic_user" query:"auth_basic_user" `
}

// HOFSTADTER_BELOW
