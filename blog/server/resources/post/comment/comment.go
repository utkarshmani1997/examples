package comment

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
Name:      comment
Route:     comments
Resource:  type.lib.types.Comment
Path:      resources.[:].resources
Parent:    post

*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var commentValidate = validator.New()

/*
Where's your docs doc?!
*/
func Handle_LIST_Comment(ctx echo.Context) (err error) {

	// HOFSTADTER_START list-body-start
	// HOFSTADTER_END   list-body-start

	// CTX path params
	postID := ctx.Param("post-id")

	// OUTPUT
	// user-defined
	// it's not a view
	var comment []types.Comment
	// fmt.Println("list comment")

	// HOFSTADTER_START list-body-mid
	// HOFSTADTER_END   list-body-mid

	// HOFSTADTER_START list-body
	// HOFSTADTER_END   list-body

	// HOFSTADTER_START list-body-end
	// HOFSTADTER_END   list-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, comment)
	return err // hacky...
}

/*
Where's your docs doc?!
*/
func Handle_POST_Comment(ctx echo.Context) (err error) {

	// HOFSTADTER_START post-body-start
	// HOFSTADTER_END   post-body-start

	// CTX path params
	postID := ctx.Param("post-id")

	// input
	// START binding input to query/form/body params
	// Initialize
	// user-defined
	// it's not a view
	var inComment types.Comment
	err = ctx.Bind(&inComment)
	if err != nil {
		return err
	}
	err = commentValidate.Struct(inComment)
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
	var outComment types.Comment
	// fmt.Println("post comment")

	// HOFSTADTER_START post-body-mid
	// HOFSTADTER_END   post-body-mid

	// HOFSTADTER_START post-body
	// HOFSTADTER_END   post-body

	// HOFSTADTER_START post-body-end
	// HOFSTADTER_END   post-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, outComment)
	return err // hacky...
}

/*
Where's your docs doc?!
*/
func Handle_GET_Comment(ctx echo.Context) (err error) {

	// HOFSTADTER_START get-body-start
	// HOFSTADTER_END   get-body-start

	// M path params
	commentID := ctx.Param("comment-id")

	// CTX path params
	postID := ctx.Param("post-id")

	// OUTPUT
	// user-defined
	// it's not a view
	var comment types.Comment
	// fmt.Println("get comment")

	// HOFSTADTER_START get-body-mid
	// HOFSTADTER_END   get-body-mid

	// hello world

	// HOFSTADTER_START get-body-end
	// HOFSTADTER_END   get-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, comment)
	return err // hacky...
}

/*
Where's your docs doc?!
*/
func Handle_DELETE_Comment(ctx echo.Context) (err error) {

	// HOFSTADTER_START delete-body-start
	// HOFSTADTER_END   delete-body-start

	// M path params
	commentID := ctx.Param("comment-id")

	// CTX path params
	postID := ctx.Param("post-id")

	// OUTPUT
	// user-defined
	// it's not a view
	var comment types.Comment
	// fmt.Println("delete comment")

	// HOFSTADTER_START delete-body-mid
	// HOFSTADTER_END   delete-body-mid

	// HOFSTADTER_START delete-body
	// HOFSTADTER_END   delete-body

	// HOFSTADTER_START delete-body-end
	// HOFSTADTER_END   delete-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, comment)
	return err // hacky...
}

/*
Where's your docs doc?!
*/
func Handle_PUT_Comment(ctx echo.Context) (err error) {

	// HOFSTADTER_START put-body-start
	// HOFSTADTER_END   put-body-start

	// M path params
	commentID := ctx.Param("comment-id")

	// CTX path params
	postID := ctx.Param("post-id")

	// input
	// START binding input to query/form/body params
	// Initialize
	// user-defined
	// it's not a view
	var inComment types.Comment
	err = ctx.Bind(&inComment)
	if err != nil {
		return err
	}
	err = commentValidate.Struct(inComment)
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
	var outComment types.Comment
	// fmt.Println("put comment")

	// HOFSTADTER_START put-body-mid
	// HOFSTADTER_END   put-body-mid

	// HOFSTADTER_START put-body
	// HOFSTADTER_END   put-body

	// HOFSTADTER_START put-body-end
	// HOFSTADTER_END   put-body-end

	// return the output response
	// should check accept-type here
	return ctx.JSON(http.StatusOK, outComment)
	return err // hacky...
}

// End resource.methods

// end resource.routes

// HOFSTADTER_BELOW
