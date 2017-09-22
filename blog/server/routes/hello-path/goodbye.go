package helloPath

import (
	"github.com/labstack/echo"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
	// custom imports
)

/*
API:     server
Name:    goodbye
Route:   goodbye
Method:  get
Path:    routes.[:].routes
Parent:  hello-path
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var goodbyeValidate = validator.New()

/*
Where's your docs doc?!
*/
func Handle_GET_Goodbye(ctx echo.Context) (err error) {
	// Check params

	// path params
	message := ctx.Param("message")

	name := ctx.Param("name")

	// OUTPUT
	// builtin
	var output string

	output = "goodbye " + name + "  " + message

	// should check accept-type here
	return ctx.String(http.StatusOK, output)

	return nil
}

// HOFSTADTER_BELOW
