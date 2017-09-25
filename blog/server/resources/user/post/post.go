package post

import (
	"github.com/pkg/errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"

	"github.com/hofstadter-io/examples/blog/lib/types"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
API:       server
Name:      post
Route:     post
Resource:  type.lib.types.Post
Path:      resources.[:].resources
Parent:    user

*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

var postValidate = validator.New()

/*
Where's your docs doc?!
*/
func Handle_LIST_Post(ctx echo.Context) (err error) {

	// HOFSTADTER_START list-body-start
	// HOFSTADTER_END   list-body-start

	// OUTPUT
	// user-defined
	// it's not a view
	var post []types.Post
	// fmt.Println("list post")

	// HOFSTADTER_START list-body-mid
	// HOFSTADTER_END   list-body-mid

	// INSERT default LIST body here based on method type

	// HOFSTADTER_START list-body-end
	// HOFSTADTER_END   list-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, post)
	return err // hacky...
}

/*
Where's your docs doc?!
*/
func Handle_POST_Post(ctx echo.Context) (err error) {

	// HOFSTADTER_START post-body-start
	// HOFSTADTER_END   post-body-start

	// input
	// START binding input to query/form/body params
	// Initialize
	// user-defined
	// it's not a view
	var inPost types.Post
	err = ctx.Bind(&inPost)
	if err != nil {
		return err
	}
	err = postValidate.Struct(inPost)
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
	// it's not a view
	var outPost types.Post
	// fmt.Println("post post")

	// HOFSTADTER_START post-body-mid
	// HOFSTADTER_END   post-body-mid

	// INSERT default POST body here based on method type

	// HOFSTADTER_START post-body-end
	// HOFSTADTER_END   post-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, outPost)
	return err // hacky...
}

/*
Where's your docs doc?!
*/
func Handle_GET_Post(ctx echo.Context) (err error) {

	// HOFSTADTER_START get-body-start
	// HOFSTADTER_END   get-body-start

	// M path params
	postUUID := ctx.Param("post-uuid")

	// OUTPUT
	// user-defined
	// it's not a view
	var post types.Post
	// fmt.Println("get post")

	// HOFSTADTER_START get-body-mid
	// HOFSTADTER_END   get-body-mid

	// INSERT default GET body here based on method type

	// HOFSTADTER_START get-body-end
	// HOFSTADTER_END   get-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, post)
	return err // hacky...
}

/*
Where's your docs doc?!
*/
func Handle_DELETE_Post(ctx echo.Context) (err error) {

	// HOFSTADTER_START delete-body-start
	// HOFSTADTER_END   delete-body-start

	// M path params
	postUUID := ctx.Param("post-uuid")

	// OUTPUT
	// user-defined
	// it's not a view
	var post types.Post
	// fmt.Println("delete post")

	// HOFSTADTER_START delete-body-mid
	// HOFSTADTER_END   delete-body-mid

	// INSERT default DELETE body here based on method type

	// HOFSTADTER_START delete-body-end
	// HOFSTADTER_END   delete-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, post)
	return err // hacky...
}

/*
Where's your docs doc?!
*/
func Handle_PUT_Post(ctx echo.Context) (err error) {

	// HOFSTADTER_START put-body-start
	// HOFSTADTER_END   put-body-start

	// M path params
	postUUID := ctx.Param("post-uuid")

	// input
	// START binding input to query/form/body params
	// Initialize
	// user-defined
	// it's not a view
	var inPost types.Post
	err = ctx.Bind(&inPost)
	if err != nil {
		return err
	}
	err = postValidate.Struct(inPost)
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
	// it's not a view
	var outPost types.Post
	// fmt.Println("put post")

	// HOFSTADTER_START put-body-mid
	// HOFSTADTER_END   put-body-mid

	// INSERT default PUT body here based on method type

	// HOFSTADTER_START put-body-end
	// HOFSTADTER_END   put-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, outPost)
	return err // hacky...
}

// End resource.methods

// end resource.routes

// HOFSTADTER_BELOW
