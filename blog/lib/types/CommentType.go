package types

import (
	"github.com/jinzhu/gorm"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
Name:      Comment
About:     The blog post comment type
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type Comment struct {

	/* ORM: server.api.databases.[name==postgres]*/

	gorm.Model
	UUID string `json:"uuid" xml:"uuid" yaml:"uuid" form:"uuid" query:"uuid" validate:"uuidv4" `

	Content string `json:"content" xml:"content" yaml:"content" form:"content" query:"content" validate:"required|alphanumunicode" `

	// Relations ----
	// Should be using some type lookup and package resolution type things

	// owned-by (the other half of has-one/many)
	PostID uint `json:"post_id" xml:"post_id" yaml:"post_id" form:"post_id" query:"post_id" `

	// belongs-to
	User User `json:"user" xml:"user" yaml:"user" form:"user" query:"user" `

	UserID uint `json:"user_id" xml:"user_id" yaml:"user_id" form:"user_id" query:"user_id" `
	// the type should be inferred from REL.type and REL.foreign-key

}

// HOFSTADTER_BELOW
