package resources

import (
	"github.com/labstack/echo"

	"github.com/hofstadter-io/examples/blog/server/resources/post"
	"github.com/hofstadter-io/examples/blog/server/resources/user"
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

	serverGroup := G.Group("")

	// HOFSTADTER_START router-start
	// HOFSTADTER_END   router-start

	// names: server | server
	// routes SAME NAME

	// resources
	err = post.InitRouter(serverGroup)
	if err != nil {
		return err
	}
	err = user.InitRouter(serverGroup)
	if err != nil {
		return err
	}

	// HOFSTADTER_START router-end
	// HOFSTADTER_END   router-end

	return nil
}

// HOFSTADTER_BELOW
