package routes

import (
	"github.com/labstack/echo"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
	// custom imports
)

/*
API:     server
Name:    hello-path
Route:   hello
Method:  get
Path:    routes
Parent:  server
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var helloPathValidate = validator.New()

/*
Where's your docs doc?!
*/
func Handle_GET_HelloPath(ctx echo.Context) (err error) {
	// Check params

	// path params
	name := ctx.Param("name")

	// OUTPUT
	// builtin
	var output string

	output = "hello " + name

	// should check accept-type here
	return ctx.String(http.StatusOK, output)

	return nil
}

// HOFSTADTER_BELOW
