package commands

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import

	// custom imports
	"fmt"
	"github.com/hofstadter-io/examples/blog/server/databases/postgres"
	"os"

	// infered imports
	// infered imports

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

var ServerToolDBLong = `A tool for working with the server databases`

var (
	RootConfigPFlag string
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&RootConfigPFlag, "config", "C", "server/config.yaml", "the config file for the API server")
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))

}

var (
	RootCmd = &cobra.Command{

		Use: "serverToolDB",

		Short: "A tool for working with the server databases",

		Long: ServerToolDBLong,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Argument Parsing

			filename := "server/config.yaml"
			f, err := os.Open(filename)
			if err != nil {
				fmt.Println("Error", err)
				os.Exit(1)
			}
			viper.SetConfigType("yaml")
			verr := viper.MergeConfig(f)
			if verr != nil {
				fmt.Println("Error", verr)
				os.Exit(1)
			}
			postgresHost := viper.GetString("databases.postgres.host")
			postgresPort := viper.GetString("databases.postgres.port")
			postgresUser := viper.GetString("databases.postgres.user")
			postgresPass := viper.GetString("databases.postgres.pass")
			postgresDb := viper.GetString("databases.postgres.dbname")
			postgresSslmode := viper.GetString("databases.postgres.sslmode")

			connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
				postgresHost,
				postgresPort,
				postgresUser,
				postgresPass,
				postgresDb,
				postgresSslmode,
			)
			postgres.ConnectToPsql(connStr)

		},

		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			logger.Debug("In PersistentPostRun serverToolDBCmd", "args", args)
			// Argument Parsing

			postgres.DisconnectFromPsql()

		},
	}
)

/*
commands:
- args:
  - ctx_path: dsl.server.tools.db-ops.cli.commands.[0].args.[0]
    name: query
    parent: serverToolDB.query
    parent_path: dsl.server.tools.db-ops.cli.commands.[0]
    pkg_path: server/tools/db-ops/cli/commands/[0]/args
    pkgPath: serverToolDB/query/query
    type: string
  ctx_path: dsl.server.tools.db-ops.cli.commands.[0]
  name: query
  parent: serverToolDB
  parent_path: dsl.server.tools.db-ops.cli
  path: commands
  pkg_path: server/tools/db-ops/cli/commands
  pkgPath: serverToolDB/query
  short: query records
- body: |
    types.MigratePostTable(postgres.POSTGRES)

    types.MigrateUserTable(postgres.POSTGRES)
  ctx_path: dsl.server.tools.db-ops.cli.commands.[1]
  imports:
  - path: github.com/hofstadter-io/examples/blog/server/databases/postgres
  - path: github.com/hofstadter-io/examples/blog/lib/types
  - path: github.com/hofstadter-io/examples/blog/lib/types
  name: migrate-tables
  parent: serverToolDB
  parent_path: dsl.server.tools.db-ops.cli
  path: commands
  pkg_path: server/tools/db-ops/cli/commands
  pkgPath: serverToolDB/migrate-tables
  short: create or migrate all of the tables
- commands:
  - body: ""
    ctx_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[0]
    imports:
    - path: github.com/hofstadter-io/examples/blog/lib/types
    - path: github.com/hofstadter-io/examples/blog/server/databases/postgres
    name: migrate
    parent: comment
    parent_path: dsl.server.tools.db-ops.cli.commands.[2]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[2]/commands
    pkgPath: serverToolDB/comment/migrate
    usage: migrate
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[1].args.[0]
      name: data-file
      parent: serverToolDB.comment.create
      parent_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[1]
      pkg_path: server/tools/db-ops/cli/commands/[2]/commands/[1]/args
      pkgPath: serverToolDB/comment/create/data-file
      required: true
      type: string
    body: |
      var comment types.Comment
      // unmarshal data file into struct

      err := types.CreateComment(postgres.POSTGRES, comment)
      if err != nil {
          fmt.Println("Error", err)
          os.Exit(1)
      }
      fmt.Println("Created:", comment.uuid)
    ctx_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[1]
    imports:
    - path: github.com/hofstadter-io/examples/blog/lib/types
    - path: github.com/hofstadter-io/examples/blog/server/databases/postgres
    name: create
    parent: comment
    parent_path: dsl.server.tools.db-ops.cli.commands.[2]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[2]/commands
    pkgPath: serverToolDB/comment/create
    usage: create <data-file>
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[2].args.[0]
      name: uuid
      parent: serverToolDB.comment.find
      parent_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[2]
      pkg_path: server/tools/db-ops/cli/commands/[2]/commands/[2]/args
      pkgPath: serverToolDB/comment/find/uuid
      type: string
    ctx_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[2]
    name: find
    parent: comment
    parent_path: dsl.server.tools.db-ops.cli.commands.[2]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[2]/commands
    pkgPath: serverToolDB/comment/find
    usage: find <uuid>
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[3].args.[0]
      name: data-file
      parent: serverToolDB.comment.update
      parent_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[3]
      pkg_path: server/tools/db-ops/cli/commands/[2]/commands/[3]/args
      pkgPath: serverToolDB/comment/update/data-file
      type: string
    ctx_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[3]
    name: update
    parent: comment
    parent_path: dsl.server.tools.db-ops.cli.commands.[2]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[2]/commands
    pkgPath: serverToolDB/comment/update
    usage: update <data-file>
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[4].args.[0]
      name: uuid
      parent: serverToolDB.comment.delete
      parent_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[4]
      pkg_path: server/tools/db-ops/cli/commands/[2]/commands/[4]/args
      pkgPath: serverToolDB/comment/delete/uuid
      type: string
    ctx_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[4]
    name: delete
    parent: comment
    parent_path: dsl.server.tools.db-ops.cli.commands.[2]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[2]/commands
    pkgPath: serverToolDB/comment/delete
    usage: delete <uuid>
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[5].args.[0]
      name: matchers
      parent: serverToolDB.comment.query
      parent_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[5]
      pkg_path: server/tools/db-ops/cli/commands/[2]/commands/[5]/args
      pkgPath: serverToolDB/comment/query/matchers
      rest: true
      type: array:string
    ctx_path: dsl.server.tools.db-ops.cli.commands.[2].commands.[5]
    name: query
    parent: comment
    parent_path: dsl.server.tools.db-ops.cli.commands.[2]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[2]/commands
    pkgPath: serverToolDB/comment/query
    usage: query [key=value]...
  ctx_path: dsl.server.tools.db-ops.cli.commands.[2]
  name: comment
  omit-run: true
  parent: serverToolDB
  parent_path: dsl.server.tools.db-ops.cli
  path: commands
  pkg_path: server/tools/db-ops/cli/commands
  pkgPath: serverToolDB/comment
  short: work with the comment resource
- commands:
  - body: ""
    ctx_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[0]
    imports:
    - path: github.com/hofstadter-io/examples/blog/lib/types
    - path: github.com/hofstadter-io/examples/blog/server/databases/postgres
    name: migrate
    parent: post
    parent_path: dsl.server.tools.db-ops.cli.commands.[3]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[3]/commands
    pkgPath: serverToolDB/post/migrate
    usage: migrate
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[1].args.[0]
      name: data-file
      parent: serverToolDB.post.create
      parent_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[1]
      pkg_path: server/tools/db-ops/cli/commands/[3]/commands/[1]/args
      pkgPath: serverToolDB/post/create/data-file
      required: true
      type: string
    body: |
      var post types.Post
      // unmarshal data file into struct

      err := types.CreatePost(postgres.POSTGRES, post)
      if err != nil {
          fmt.Println("Error", err)
          os.Exit(1)
      }
      fmt.Println("Created:", post.uuid)
    ctx_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[1]
    imports:
    - path: github.com/hofstadter-io/examples/blog/lib/types
    - path: github.com/hofstadter-io/examples/blog/server/databases/postgres
    name: create
    parent: post
    parent_path: dsl.server.tools.db-ops.cli.commands.[3]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[3]/commands
    pkgPath: serverToolDB/post/create
    usage: create <data-file>
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[2].args.[0]
      name: uuid
      parent: serverToolDB.post.find
      parent_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[2]
      pkg_path: server/tools/db-ops/cli/commands/[3]/commands/[2]/args
      pkgPath: serverToolDB/post/find/uuid
      type: string
    ctx_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[2]
    name: find
    parent: post
    parent_path: dsl.server.tools.db-ops.cli.commands.[3]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[3]/commands
    pkgPath: serverToolDB/post/find
    usage: find <uuid>
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[3].args.[0]
      name: data-file
      parent: serverToolDB.post.update
      parent_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[3]
      pkg_path: server/tools/db-ops/cli/commands/[3]/commands/[3]/args
      pkgPath: serverToolDB/post/update/data-file
      type: string
    ctx_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[3]
    name: update
    parent: post
    parent_path: dsl.server.tools.db-ops.cli.commands.[3]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[3]/commands
    pkgPath: serverToolDB/post/update
    usage: update <data-file>
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[4].args.[0]
      name: uuid
      parent: serverToolDB.post.delete
      parent_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[4]
      pkg_path: server/tools/db-ops/cli/commands/[3]/commands/[4]/args
      pkgPath: serverToolDB/post/delete/uuid
      type: string
    ctx_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[4]
    name: delete
    parent: post
    parent_path: dsl.server.tools.db-ops.cli.commands.[3]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[3]/commands
    pkgPath: serverToolDB/post/delete
    usage: delete <uuid>
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[5].args.[0]
      name: matchers
      parent: serverToolDB.post.query
      parent_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[5]
      pkg_path: server/tools/db-ops/cli/commands/[3]/commands/[5]/args
      pkgPath: serverToolDB/post/query/matchers
      rest: true
      type: array:string
    ctx_path: dsl.server.tools.db-ops.cli.commands.[3].commands.[5]
    name: query
    parent: post
    parent_path: dsl.server.tools.db-ops.cli.commands.[3]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[3]/commands
    pkgPath: serverToolDB/post/query
    usage: query [key=value]...
  ctx_path: dsl.server.tools.db-ops.cli.commands.[3]
  name: post
  omit-run: true
  parent: serverToolDB
  parent_path: dsl.server.tools.db-ops.cli
  path: commands
  pkg_path: server/tools/db-ops/cli/commands
  pkgPath: serverToolDB/post
  short: work with the post resource
- commands:
  - body: ""
    ctx_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[0]
    imports:
    - path: github.com/hofstadter-io/examples/blog/lib/types
    - path: github.com/hofstadter-io/examples/blog/server/databases/postgres
    name: migrate
    parent: user
    parent_path: dsl.server.tools.db-ops.cli.commands.[4]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[4]/commands
    pkgPath: serverToolDB/user/migrate
    usage: migrate
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[1].args.[0]
      name: data-file
      parent: serverToolDB.user.create
      parent_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[1]
      pkg_path: server/tools/db-ops/cli/commands/[4]/commands/[1]/args
      pkgPath: serverToolDB/user/create/data-file
      required: true
      type: string
    body: |
      var user types.User
      // unmarshal data file into struct

      err := types.CreateUser(postgres.POSTGRES, user)
      if err != nil {
          fmt.Println("Error", err)
          os.Exit(1)
      }
      fmt.Println("Created:", user.uuid)
    ctx_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[1]
    imports:
    - path: github.com/hofstadter-io/examples/blog/lib/types
    - path: github.com/hofstadter-io/examples/blog/server/databases/postgres
    name: create
    parent: user
    parent_path: dsl.server.tools.db-ops.cli.commands.[4]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[4]/commands
    pkgPath: serverToolDB/user/create
    usage: create <data-file>
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[2].args.[0]
      name: uuid
      parent: serverToolDB.user.find
      parent_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[2]
      pkg_path: server/tools/db-ops/cli/commands/[4]/commands/[2]/args
      pkgPath: serverToolDB/user/find/uuid
      type: string
    ctx_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[2]
    name: find
    parent: user
    parent_path: dsl.server.tools.db-ops.cli.commands.[4]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[4]/commands
    pkgPath: serverToolDB/user/find
    usage: find <uuid>
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[3].args.[0]
      name: data-file
      parent: serverToolDB.user.update
      parent_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[3]
      pkg_path: server/tools/db-ops/cli/commands/[4]/commands/[3]/args
      pkgPath: serverToolDB/user/update/data-file
      type: string
    ctx_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[3]
    name: update
    parent: user
    parent_path: dsl.server.tools.db-ops.cli.commands.[4]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[4]/commands
    pkgPath: serverToolDB/user/update
    usage: update <data-file>
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[4].args.[0]
      name: uuid
      parent: serverToolDB.user.delete
      parent_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[4]
      pkg_path: server/tools/db-ops/cli/commands/[4]/commands/[4]/args
      pkgPath: serverToolDB/user/delete/uuid
      type: string
    ctx_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[4]
    name: delete
    parent: user
    parent_path: dsl.server.tools.db-ops.cli.commands.[4]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[4]/commands
    pkgPath: serverToolDB/user/delete
    usage: delete <uuid>
  - args:
    - ctx_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[5].args.[0]
      name: matchers
      parent: serverToolDB.user.query
      parent_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[5]
      pkg_path: server/tools/db-ops/cli/commands/[4]/commands/[5]/args
      pkgPath: serverToolDB/user/query/matchers
      rest: true
      type: array:string
    ctx_path: dsl.server.tools.db-ops.cli.commands.[4].commands.[5]
    name: query
    parent: user
    parent_path: dsl.server.tools.db-ops.cli.commands.[4]
    path: commands.[:].commands
    pkg_path: server/tools/db-ops/cli/commands/[4]/commands
    pkgPath: serverToolDB/user/query
    usage: query [key=value]...
  ctx_path: dsl.server.tools.db-ops.cli.commands.[4]
  name: user
  omit-run: true
  parent: serverToolDB
  parent_path: dsl.server.tools.db-ops.cli
  path: commands
  pkg_path: server/tools/db-ops/cli/commands
  pkgPath: serverToolDB/user
  short: work with the user resource
ctx_path: dsl.server.tools.db-ops.cli
imports:
- path: fmt
- path: os
- path: github.com/hofstadter-io/examples/blog/server/databases/postgres
long: A tool for working with the server databases
name: serverToolDB
omit-run: true
parent: ""
parent_path: ""
persistent-postrun: true
persistent-postrun-body: |+
  postgres.DisconnectFromPsql()

persistent-prerun: true
persistent-prerun-body: |+
  filename := "server/config.yaml"
  f, err := os.Open(filename)
  if err != nil {
      fmt.Println("Error", err)
      os.Exit(1)
  }
  viper.SetConfigType("yaml")
  verr := viper.MergeConfig(f)
  if verr != nil {
      fmt.Println("Error", verr)
      os.Exit(1)
  }
  postgresHost := viper.GetString("databases.postgres.host")
  postgresPort := viper.GetString("databases.postgres.port")
  postgresUser := viper.GetString("databases.postgres.user")
  postgresPass := viper.GetString("databases.postgres.pass")
  postgresDb := viper.GetString("databases.postgres.dbname")
  postgresSslmode := viper.GetString("databases.postgres.sslmode")

  connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
      postgresHost,
      postgresPort,
      postgresUser,
      postgresPass,
      postgresDb,
      postgresSslmode,
  )
  postgres.ConnectToPsql(connStr)

pflags:
- ctx_path: dsl.server.tools.db-ops.cli.pflags.[0]
  default: server/config.yaml
  help: the config file for the API server
  long: config
  name: config
  parent: serverToolDB
  parent_path: dsl.server.tools.db-ops.cli
  pkg_path: server/tools/db-ops/cli/pflags
  pkgPath: serverToolDB/config
  short: C
  type: string
pkg_path: server/tools/db-ops
pkgPath: serverToolDB
relPath: server/tools/db-ops/commands
short: A tool for working with the server databases

*/

// HOFSTADTER_BELOW
