package routes

import (
	"github.com/labstack/echo"

	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports
	"github.com/go-playground/validator"
	"github.com/hofstadter-io/examples/blog/lib/types"
)

/*
API:     server
Name:    Signup
Route:   signup
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

var SignupValidate = validator.New()

/*
Where's your docs doc?!
*/
func Handle_POST_Signup(ctx echo.Context) (err error) {
	// Check params

	// input
	// START binding input to query/form/body params
	// Initialize
	// user-defined
	// it's a view!
	var input types.AuthBasicUserSignupRequest
	err = ctx.Bind(&input)
	if err != nil {
		return err
	}
	err = SignupValidate.Struct(input)
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
	var output types.AuthBasicUserSignupResponse

	// hello signup!
	/*
	   auth-basic

	*/

	// should check accept-type here
	return ctx.JSON(http.StatusOK, output)

	return nil
}

// HOFSTADTER_BELOW
