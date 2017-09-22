package user

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

	userGroup := G.Group("/user")

	// HOFSTADTER_START router-start
	// HOFSTADTER_END   router-start

	// names: server | user
	// routes NOT SAME NAME

	// methods
	userGroup.GET("", Handle_LIST_User)
	userGroup.POST("", Handle_POST_User)
	userGroup.GET("/:user-id", Handle_GET_User)
	userGroup.DELETE("/:user-id", Handle_DELETE_User)
	userGroup.PUT("/:user-id", Handle_PUT_User)

	// HOFSTADTER_START router-end
	// HOFSTADTER_END   router-end

	return nil
}

// HOFSTADTER_BELOW
