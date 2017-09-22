package helloPath

import (
	"github.com/labstack/echo"

	"github.com/hofstadter-io/examples/blog/server/routes/hello-path"
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

	helloPathGroup := G.Group("/hello")

	// HOFSTADTER_START router-start
	// HOFSTADTER_END   router-start

	// names: server | hello-path
	// routes NOT SAME NAME

	// routes RESOURCE
	helloPathGroup.GET("/goodbye/:message", Handle_GET_Goodbye)

	// HOFSTADTER_START router-end
	// HOFSTADTER_END   router-end

	return nil
}

// HOFSTADTER_BELOW
