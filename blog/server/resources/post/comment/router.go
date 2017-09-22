package comment

import (
	"github.com/labstack/echo"
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

	commentGroup := G.Group("/comments")

	// HOFSTADTER_START router-start
	// HOFSTADTER_END   router-start

	// names: server | comment
	// routes NOT SAME NAME

	// methods
	commentsGroup.GET("", Handle_LIST_Comment)
	commentsGroup.POST("", Handle_POST_Comment)
	commentsGroup.GET("/:comment-id", Handle_GET_Comment)
	commentsGroup.DELETE("/:comment-id", Handle_DELETE_Comment)
	commentsGroup.PUT("/:comment-id", Handle_PUT_Comment)

	// HOFSTADTER_START router-end
	// HOFSTADTER_END   router-end

	return nil
}

// HOFSTADTER_BELOW
