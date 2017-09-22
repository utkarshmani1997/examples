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

// Should find a way to build up errors and return all
// POST  map[pkgPath:server/Signup/input name:input type:type.lib.types.User.views.AuthBasicUserSignupRequest parent:server.Signup parent_path:dsl.server.api.routes.[1] ctx_path:dsl.server.api.routes.[1].input pkg_path:server/api/routes/[1]]  ->  map[pkg_path:server/api/routes/[1] pkgPath:server/Signup/output name:output type:type.lib.types.User.views.AuthBasicUserSignupResponse parent:server.Signup parent_path:dsl.server.api.routes.[1] ctx_path:dsl.server.api.routes.[1].output]
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
	/*
			ctx_path: dsl.server.api.routes.[1].output
	name: output
	parent: server.Signup
	parent_path: dsl.server.api.routes.[1]
	pkg_path: server/api/routes/[1]
	pkgPath: server/Signup/output
	type: type.lib.types.User.views.AuthBasicUserSignupResponse

	*/
	// user-defined
	// it's a view!
	var output types.AuthBasicUserSignupResponse

	// BODY: something needs to go here

	// should check accept-type here
	return ctx.JSON(http.StatusOK, output)

	return nil
}

// HOFSTADTER_BELOW
