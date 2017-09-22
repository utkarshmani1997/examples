package post

import (
	"github.com/labstack/echo"

	"github.com/hofstadter-io/examples/blog/server/resources/post/comment"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

func InitRouter(G *echo.Group) (err error) {

	// HOFSTADTER_START router-pre
	// HOFSTADTER_END   router-pre

	postGroup := G.Group("/post")

	// HOFSTADTER_START router-start
	// HOFSTADTER_END   router-start

	// names: server | post
	// routes NOT SAME NAME

	// methods
	postGroup.GET("", Handle_LIST_Post)
	postGroup.POST("", Handle_POST_Post)
	postGroup.GET("/:post-id", Handle_GET_Post)
	postGroup.DELETE("/:post-id", Handle_DELETE_Post)
	postGroup.PUT("/:post-id", Handle_PUT_Post)

	// resources
	err = comment.InitRouter(postGroup)
	if err != nil {
		return err
	}

	// HOFSTADTER_START router-end
	// HOFSTADTER_END   router-end

	return nil
}

// HOFSTADTER_BELOW
