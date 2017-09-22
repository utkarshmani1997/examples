package routes

import (
	"github.com/labstack/echo"

	"github.com/hofstadter-io/examples/blog/server/routes"
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

	serverGroup.GET("/hello/:name", Handle_GET_HelloPath)

	serverGroup.POST("/signup", Handle_POST_Signup)

	serverGroup.POST("/login", Handle_POST_Login)

	addPrometheusHandlers(serverGroup)

	addKubernetesHandlers(serverGroup)

	// HOFSTADTER_START router-end
	// HOFSTADTER_END   router-end

	return nil
}

// HOFSTADTER_BELOW
