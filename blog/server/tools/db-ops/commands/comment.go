package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports

	// infered imports

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/examples/blog/server/tools/db-ops/commands/comment"
)

// Tool:   serverToolDB
// Name:   comment
// Usage:
// Parent: serverToolDB

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var CommentCmd = &cobra.Command{

	Use: "comment",

	Short: "work with the comment resource",
}

func init() {
	RootCmd.AddCommand(CommentCmd)
}

func init() {
	// add sub-commands to this command when present

	CommentCmd.AddCommand(comment.MigrateCmd)
	CommentCmd.AddCommand(comment.CreateCmd)
	CommentCmd.AddCommand(comment.FindCmd)
	CommentCmd.AddCommand(comment.UpdateCmd)
	CommentCmd.AddCommand(comment.DeleteCmd)
	CommentCmd.AddCommand(comment.QueryCmd)
}

// HOFSTADTER_BELOW
