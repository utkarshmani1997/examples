package types

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      Post
About:     The blog post type
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

func NewPost() *Post {
	return &Post{
		PublishTime: nil,
	}
}

/*
Where's your docs doc?!
*/
func (P *Post) Publish(doPub bool) (err error) {
	// HOFSTADTER_START Publish

	// HOFSTADTER_END   Publish
	return
}

// HOFSTADTER_BELOW
