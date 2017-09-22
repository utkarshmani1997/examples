package routes

import (
	"github.com/labstack/echo"

	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports
	"github.com/go-playground/validator"
)

/*
API:     server
Name:    Login
Route:   login
Method:  POST
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

var LoginValidate = validator.New()

/*
Where's your docs doc?!
*/
func Handle_POST_Login(ctx echo.Context) (err error) {
	// Check params

	// input
	// START binding input to query/form/body params
	// Initialize
	// user-defined
	// it's a view!
	var input types.AuthBasicUserLoginRequest
	err = ctx.Bind(&input)
	if err != nil {
		return err
	}
	err = LoginValidate.Struct(input)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"errors": err,
				"type":   "invalid",
			})
		}
		if _, ok := err.(*validator.ValidationErrors); ok {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"errors": err,
				"type":   "validation",
			})
		}
		return err
	}
	// END binding input to query/form/body params

	// OUTPUT
	// user-defined
	// it's a view!
	var output types.AuthBasicUserLoginResponse

	// theBODY: something needs to go here

	// should check accept-type here
	return ctx.JSON(http.StatusOK, output)

	return nil
}

// HOFSTADTER_BELOW
