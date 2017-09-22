package types

import (
	"time"

	"github.com/jinzhu/gorm"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
Name:      Post
About:     The blog post type
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type Post struct {

	/* ORM: server.api.databases.[name==postgres]*/

	gorm.Model
	UUID string `json:"uuid" xml:"uuid" yaml:"uuid" form:"uuid" query:"uuid" validate:"uuidv4" `

	Title string `json:"title" xml:"title" yaml:"title" form:"title" query:"title" validate:"required|alphanumunicode|min=1|max=256" `

	Content string `json:"content" xml:"content" yaml:"content" form:"content" query:"content" validate:"alphanumunicode" `

	Draft bool `json:"draft" xml:"draft" yaml:"draft" form:"draft" query:"draft" `

	PublishTime *time.Time `json:"publish_time" xml:"publish_time" yaml:"publish_time" form:"publish_time" query:"publish_time" `

	// Relations ----
	// Should be using some type lookup and package resolution type things

	// owned-by (the other half of has-one/many)
	UserID uint `json:"user_id" xml:"user_id" yaml:"user_id" form:"user_id" query:"user_id" `

	// has-many
	Comments []Comment `json:"comments" xml:"comments" yaml:"comments" form:"comments" query:"comments" `

	// many-to-many
	Likes []User `json:"likes" xml:"likes" yaml:"likes" form:"likes" query:"likes" gorm:"many2many:user_likes"`
}

// HOFSTADTER_BELOW
