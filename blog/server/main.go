package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"
	log15 "gopkg.in/inconshreveable/log15.v2"

	"github.com/hofstadter-io/examples/blog/server/resources"

	"github.com/hofstadter-io/examples/blog/server/routes"

	"github.com/hofstadter-io/examples/blog/server/databases/postgres"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// Name:     server
// Version:
// About:    a simple blog server

// HOFSTADTER_START start
// HOFSTADTER_END   start

func main() {
	var err error

	// load the configuration file
	read_config()

	// configure the logger
	config_logger()

	// connect to the database(s)
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
	defer postgres.DisconnectFromPsql()

	// create the echo server object
	E := echo.New()

	// HOFSTADTER_START main-pre-group
	// HOFSTADTER_END   main-pre-group

	// Base API Group
	G := E.Group("api/v1")

	// HOFSTADTER_START main-pre-middleware
	// HOFSTADTER_END   main-pre-middleware

	// Add Middleware
	err = AddMiddleware(G, []string{"default"})
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// E.Use(middleware.Recover())
	// HOFSTADTER_START main-pre-routes
	// HOFSTADTER_END   main-pre-routes

	err = resources.InitRouter(G)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	err = routes.InitRouter(G)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	appHost := viper.GetString("host")
	appPort := viper.GetString("port")

	// HOFSTADTER_START main-prerun
	// HOFSTADTER_END   main-prerun

	E.Logger.SetLevel(log.INFO)
	E.Logger.Fatal(E.Start(appHost + ":" + appPort))
}

func read_config() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.MergeInConfig()

	// Hackery because viper only takes the first config file found... not merging, wtf does merge config mean then anyway
	f, err := os.Open("config.yml")
	if err != nil {
		f = nil
		f2, err2 := os.Open("config.yaml")
		if err2 != nil {
			f = nil
		} else {
			f = f2
		}
	}
	if f != nil {
		verr := viper.MergeConfig(f)
		if verr != nil {
			panic(verr)
		}
	} else {
		panic("missing config.yaml during start up")
	}
}

func config_logger() {
	// log-config default global values
	level := log15.LvlWarn
	stack := false

	// look up in config
	lcfg := viper.GetStringMap("log-config.default")

	if lcfg != nil && len(lcfg) > 0 {
		level_str := lcfg["level"].(string)
		stack = lcfg["stack"].(bool)
		level_local, err := log15.LvlFromString(level_str)
		if err != nil {
			panic(err)
		}
		level = level_local
	}

	termlog := log15.LvlFilterHandler(level, log15.StdoutHandler)
	if stack {
		term_stack := log15.CallerStackHandler("%+v", log15.StdoutHandler)
		termlog = log15.LvlFilterHandler(level, term_stack)
	}

	logger.SetHandler(termlog)

	// set package loggers
	resources.SetLogger(logger)

	routes.SetLogger(logger)

	// HOFSTADTER_START config-logger
	// HOFSTADTER_END   config-logger
}

/*

about: a simple blog server
auth:
  authentication:
  - type: basic
  - type: apikey
  authorization:
  - roles:
    - admin
    - user
    type: roles
  parent: server
  path: auth
client-tools: true
config:
  base-url: api/v1
  host: localhost
  port: 1248
ctx_path: dsl.server.api
data-tools: true
database-tools: true
databases:
- config:
    dbname: develop
    host: localhost
    pass: password
    port: 5432
    postgres-password: postgres
    user: admin
  ctx_path: dsl.server.api.databases.[0]
  name: postgres
  parent: server
  parent_path: dsl.server.api
  path: server.api.databases
  pkg_path: server/api/databases
  pkgPath: server/postgres
  type: psql
global-params:
- ctx_path: dsl.server.api.global-params.[0]
  name: version
  parent: server
  parent_path: dsl.server.api
  pkg_path: server/api/global-params
  pkgPath: server/version
  required: true
  type: string
kubernetes-endpoints: true
middleware:
- ctx_path: dsl.server.api.middleware.[0]
  name: auth
  parent: server
  parent_path: dsl.server.api
  pkg_path: server/api/middleware
  pkgPath: server/auth
- ctx_path: dsl.server.api.middleware.[1]
  name: cors
  parent: server
  parent_path: dsl.server.api
  pkg_path: server/api/middleware
  pkgPath: server/cors
- ctx_path: dsl.server.api.middleware.[2]
  name: csrf
  parent: server
  parent_path: dsl.server.api
  pkg_path: server/api/middleware
  pkgPath: server/csrf
- ctx_path: dsl.server.api.middleware.[3]
  name: limiter
  parent: server
  parent_path: dsl.server.api
  pkg_path: server/api/middleware
  pkgPath: server/limiter
name: server
parent: ""
parent_path: ""
pkg_path: server
pkgPath: server
prometheus-endpoints: true
relPath: server/routes
resources:
- ctx_path: dsl.server.api.resources.[0]
  methods:
  - method: list
    output:
      ctx_path: dsl.server.api.resources.[0].methods.[0].output
      name: post
      parent: server.post
      parent_path: dsl.server.api.resources.[0]
      pkg_path: server/api/resources/[0]/methods/[0]
      pkgPath: server/post/post
      type: array:type.lib.types.Post
  - input:
      ctx_path: dsl.server.api.resources.[0].methods.[1].input
      name: inPost
      parent: server.post
      parent_path: dsl.server.api.resources.[0]
      pkg_path: server/api/resources/[0]/methods/[1]
      pkgPath: server/post/inPost
      type: type.lib.types.Post
    method: post
    output:
      ctx_path: dsl.server.api.resources.[0].methods.[1].output
      name: outPost
      parent: server.post
      parent_path: dsl.server.api.resources.[0]
      pkg_path: server/api/resources/[0]/methods/[1]
      pkgPath: server/post/outPost
      type: type.lib.types.Post
  - body: // hello world
    method: get
    output:
      ctx_path: dsl.server.api.resources.[0].methods.[2].output
      name: post
      parent: server.post
      parent_path: dsl.server.api.resources.[0]
      pkg_path: server/api/resources/[0]/methods/[2]
      pkgPath: server/post/post
      type: type.lib.types.Post
    path-params:
    - ctx_path: dsl.server.api.resources.[0].methods.[2].path-params.[0]
      name: post-id
      parent: server.post
      parent_path: dsl.server.api.resources.[0]
      pkg_path: server/api/resources/[0]/methods/[2]/path-params
      pkgPath: server/post/post-id
      type: type.lib.types.Post.fields.uuid
  - method: delete
    output:
      ctx_path: dsl.server.api.resources.[0].methods.[3].output
      name: post
      parent: server.post
      parent_path: dsl.server.api.resources.[0]
      pkg_path: server/api/resources/[0]/methods/[3]
      pkgPath: server/post/post
      type: type.lib.types.Post
    path-params:
    - ctx_path: dsl.server.api.resources.[0].methods.[3].path-params.[0]
      name: post-id
      parent: server.post
      parent_path: dsl.server.api.resources.[0]
      pkg_path: server/api/resources/[0]/methods/[3]/path-params
      pkgPath: server/post/post-id
      type: type.lib.types.Post.fields.uuid
  - input:
      ctx_path: dsl.server.api.resources.[0].methods.[4].input
      name: in-post
      parent: server.post
      parent_path: dsl.server.api.resources.[0]
      pkg_path: server/api/resources/[0]/methods/[4]
      pkgPath: server/post/in-post
      type: type.lib.types.Post
    method: put
    output:
      ctx_path: dsl.server.api.resources.[0].methods.[4].output
      name: out-post
      parent: server.post
      parent_path: dsl.server.api.resources.[0]
      pkg_path: server/api/resources/[0]/methods/[4]
      pkgPath: server/post/out-post
      type: type.lib.types.Post
    path-params:
    - ctx_path: dsl.server.api.resources.[0].methods.[4].path-params.[0]
      name: post-id
      parent: server.post
      parent_path: dsl.server.api.resources.[0]
      pkg_path: server/api/resources/[0]/methods/[4]/path-params
      pkgPath: server/post/post-id
      type: type.lib.types.Post.fields.uuid
  name: post
  parent: server
  parent_path: dsl.server.api
  path: resources
  pkg_path: server/api/resources
  pkgPath: server/post
  resource: type.lib.types.Post
  resources:
  - ctx_path: dsl.server.api.resources.[0].resources.[0]
    methods:
    - method: list
      output:
        ctx_path: dsl.server.api.resources.[0].resources.[0].methods.[0].output
        name: comment
        parent: server.post.comment
        parent_path: dsl.server.api.resources.[0].resources.[0]
        pkg_path: server/api/resources/[0]/resources/[0]/methods/[0]
        pkgPath: server/post/comment/comment
        type: array:type.lib.types.Comment
    - input:
        ctx_path: dsl.server.api.resources.[0].resources.[0].methods.[1].input
        name: inComment
        parent: server.post.comment
        parent_path: dsl.server.api.resources.[0].resources.[0]
        pkg_path: server/api/resources/[0]/resources/[0]/methods/[1]
        pkgPath: server/post/comment/inComment
        type: type.lib.types.Comment
      method: post
      output:
        ctx_path: dsl.server.api.resources.[0].resources.[0].methods.[1].output
        name: outComment
        parent: server.post.comment
        parent_path: dsl.server.api.resources.[0].resources.[0]
        pkg_path: server/api/resources/[0]/resources/[0]/methods/[1]
        pkgPath: server/post/comment/outComment
        type: type.lib.types.Comment
    - body: // hello world
      method: get
      output:
        ctx_path: dsl.server.api.resources.[0].resources.[0].methods.[2].output
        name: comment
        parent: server.post.comment
        parent_path: dsl.server.api.resources.[0].resources.[0]
        pkg_path: server/api/resources/[0]/resources/[0]/methods/[2]
        pkgPath: server/post/comment/comment
        type: type.lib.types.Comment
      path-params:
      - ctx_path: dsl.server.api.resources.[0].resources.[0].methods.[2].path-params.[0]
        name: comment-id
        parent: server.post.comment
        parent_path: dsl.server.api.resources.[0].resources.[0]
        pkg_path: server/api/resources/[0]/resources/[0]/methods/[2]/path-params
        pkgPath: server/post/comment/comment-id
        type: type.lib.types.Comment.fields.uuid
    - method: delete
      output:
        ctx_path: dsl.server.api.resources.[0].resources.[0].methods.[3].output
        name: comment
        parent: server.post.comment
        parent_path: dsl.server.api.resources.[0].resources.[0]
        pkg_path: server/api/resources/[0]/resources/[0]/methods/[3]
        pkgPath: server/post/comment/comment
        type: type.lib.types.Comment
      path-params:
      - ctx_path: dsl.server.api.resources.[0].resources.[0].methods.[3].path-params.[0]
        name: comment-id
        parent: server.post.comment
        parent_path: dsl.server.api.resources.[0].resources.[0]
        pkg_path: server/api/resources/[0]/resources/[0]/methods/[3]/path-params
        pkgPath: server/post/comment/comment-id
        type: type.lib.types.Comment.fields.uuid
    - input:
        ctx_path: dsl.server.api.resources.[0].resources.[0].methods.[4].input
        name: in-comment
        parent: server.post.comment
        parent_path: dsl.server.api.resources.[0].resources.[0]
        pkg_path: server/api/resources/[0]/resources/[0]/methods/[4]
        pkgPath: server/post/comment/in-comment
        type: type.lib.types.Comment
      method: put
      output:
        ctx_path: dsl.server.api.resources.[0].resources.[0].methods.[4].output
        name: out-comment
        parent: server.post.comment
        parent_path: dsl.server.api.resources.[0].resources.[0]
        pkg_path: server/api/resources/[0]/resources/[0]/methods/[4]
        pkgPath: server/post/comment/out-comment
        type: type.lib.types.Comment
      path-params:
      - ctx_path: dsl.server.api.resources.[0].resources.[0].methods.[4].path-params.[0]
        name: comment-id
        parent: server.post.comment
        parent_path: dsl.server.api.resources.[0].resources.[0]
        pkg_path: server/api/resources/[0]/resources/[0]/methods/[4]/path-params
        pkgPath: server/post/comment/comment-id
        type: type.lib.types.Comment.fields.uuid
    name: comment
    parent: post
    parent_path: dsl.server.api.resources.[0]
    path: resources.[:].resources
    path-params:
    - ctx_path: dsl.server.api.resources.[0].resources.[0].path-params.[0]
      name: post-id
      parent: server.post.comment
      parent_path: dsl.server.api.resources.[0].resources.[0]
      pkg_path: server/api/resources/[0]/resources/[0]/path-params
      pkgPath: server/post/comment/post-id
      type: type.lib.types.Post.fields.uuid
    pkg_path: server/api/resources/[0]/resources
    pkgPath: server/post/comment
    resource: type.lib.types.Comment
    route: comments
  route: post
- ctx_path: dsl.server.api.resources.[1]
  methods:
  - method: list
    output:
      ctx_path: dsl.server.api.resources.[1].methods.[0].output
      name: user
      parent: server.user
      parent_path: dsl.server.api.resources.[1]
      pkg_path: server/api/resources/[1]/methods/[0]
      pkgPath: server/user/user
      type: array:type.lib.types.User
  - input:
      ctx_path: dsl.server.api.resources.[1].methods.[1].input
      name: inUser
      parent: server.user
      parent_path: dsl.server.api.resources.[1]
      pkg_path: server/api/resources/[1]/methods/[1]
      pkgPath: server/user/inUser
      type: type.lib.types.User
    method: post
    output:
      ctx_path: dsl.server.api.resources.[1].methods.[1].output
      name: outUser
      parent: server.user
      parent_path: dsl.server.api.resources.[1]
      pkg_path: server/api/resources/[1]/methods/[1]
      pkgPath: server/user/outUser
      type: type.lib.types.User
  - body: // hello world
    method: get
    output:
      ctx_path: dsl.server.api.resources.[1].methods.[2].output
      name: user
      parent: server.user
      parent_path: dsl.server.api.resources.[1]
      pkg_path: server/api/resources/[1]/methods/[2]
      pkgPath: server/user/user
      type: type.lib.types.User
    path-params:
    - ctx_path: dsl.server.api.resources.[1].methods.[2].path-params.[0]
      name: user-id
      parent: server.user
      parent_path: dsl.server.api.resources.[1]
      pkg_path: server/api/resources/[1]/methods/[2]/path-params
      pkgPath: server/user/user-id
      type: type.lib.types.User.fields.uuid
  - method: delete
    output:
      ctx_path: dsl.server.api.resources.[1].methods.[3].output
      name: user
      parent: server.user
      parent_path: dsl.server.api.resources.[1]
      pkg_path: server/api/resources/[1]/methods/[3]
      pkgPath: server/user/user
      type: type.lib.types.User
    path-params:
    - ctx_path: dsl.server.api.resources.[1].methods.[3].path-params.[0]
      name: user-id
      parent: server.user
      parent_path: dsl.server.api.resources.[1]
      pkg_path: server/api/resources/[1]/methods/[3]/path-params
      pkgPath: server/user/user-id
      type: type.lib.types.User.fields.uuid
  - input:
      ctx_path: dsl.server.api.resources.[1].methods.[4].input
      name: in-user
      parent: server.user
      parent_path: dsl.server.api.resources.[1]
      pkg_path: server/api/resources/[1]/methods/[4]
      pkgPath: server/user/in-user
      type: type.lib.types.User
    method: put
    output:
      ctx_path: dsl.server.api.resources.[1].methods.[4].output
      name: out-user
      parent: server.user
      parent_path: dsl.server.api.resources.[1]
      pkg_path: server/api/resources/[1]/methods/[4]
      pkgPath: server/user/out-user
      type: type.lib.types.User
    path-params:
    - ctx_path: dsl.server.api.resources.[1].methods.[4].path-params.[0]
      name: user-id
      parent: server.user
      parent_path: dsl.server.api.resources.[1]
      pkg_path: server/api/resources/[1]/methods/[4]/path-params
      pkgPath: server/user/user-id
      type: type.lib.types.User.fields.uuid
  name: user
  parent: server
  parent_path: dsl.server.api
  path: resources
  pkg_path: server/api/resources
  pkgPath: server/user
  resource: type.lib.types.User
  route: user
routes:
- body: output = "hello " + name
  ctx_path: dsl.server.api.routes.[0]
  method: get
  name: hello-path
  output:
    ctx_path: dsl.server.api.routes.[0].output
    name: output
    parent: server.hello-path
    parent_path: dsl.server.api.routes.[0]
    pkg_path: server/api/routes/[0]
    pkgPath: server/hello-path/output
    type: string
  parent: server
  parent_path: dsl.server.api
  path: routes
  pkg_path: server/api/routes
  pkgPath: server/hello-path
  post-path-params:
  - ctx_path: dsl.server.api.routes.[0].post-path-params.[0]
    name: name
    parent: server.hello-path
    parent_path: dsl.server.api.routes.[0]
    pkg_path: server/api/routes/[0]/post-path-params
    pkgPath: server/hello-path/name
    type: string
  route: hello
  routes:
  - body: output = "goodbye " + name + "  " + message
    ctx_path: dsl.server.api.routes.[0].routes.[0]
    method: get
    name: goodbye
    output:
      ctx_path: dsl.server.api.routes.[0].routes.[0].output
      name: output
      parent: server.hello-path.goodbye
      parent_path: dsl.server.api.routes.[0].routes.[0]
      pkg_path: server/api/routes/[0]/routes/[0]
      pkgPath: server/hello-path/goodbye/output
      type: string
    parent: hello-path
    parent_path: dsl.server.api.routes.[0]
    path: routes.[:].routes
    pkg_path: server/api/routes/[0]/routes
    pkgPath: server/hello-path/goodbye
    post-path-params:
    - ctx_path: dsl.server.api.routes.[0].routes.[0].post-path-params.[0]
      name: message
      parent: server.hello-path.goodbye
      parent_path: dsl.server.api.routes.[0].routes.[0]
      pkg_path: server/api/routes/[0]/routes/[0]/post-path-params
      pkgPath: server/hello-path/goodbye/message
      type: string
    route: goodbye
- body: '// BODY: something needs to go here'
  ctx_path: dsl.server.api.routes.[1]
  imports:
  - path: github.com/hofstadter-io/examples/blog/lib/types
  - path: github.com/go-playground/validator
  input:
    ctx_path: dsl.server.api.routes.[1].input
    name: input
    parent: server.Signup
    parent_path: dsl.server.api.routes.[1]
    pkg_path: server/api/routes/[1]
    pkgPath: server/Signup/input
    type: type.lib.types.User.views.AuthBasicUserSignupRequest
  method: POST
  name: Signup
  output:
    ctx_path: dsl.server.api.routes.[1].output
    name: output
    parent: server.Signup
    parent_path: dsl.server.api.routes.[1]
    pkg_path: server/api/routes/[1]
    pkgPath: server/Signup/output
    type: type.lib.types.User.views.AuthBasicUserSignupResponse
  parent: server
  parent_path: dsl.server.api
  path: routes
  pkg_path: server/api/routes
  pkgPath: server/Signup
  route: signup
- body: '// theBODY: something needs to go here'
  ctx_path: dsl.server.api.routes.[2]
  imports:
  - path: github.com/go-playground/validator
  input:
    ctx_path: dsl.server.api.routes.[2].input
    name: input
    parent: server.Login
    parent_path: dsl.server.api.routes.[2]
    pkg_path: server/api/routes/[2]
    pkgPath: server/Login/input
    type: type.lib.types.User.views.AuthBasicUserLoginRequest
  method: POST
  name: Login
  output:
    ctx_path: dsl.server.api.routes.[2].output
    name: output
    parent: server.Login
    parent_path: dsl.server.api.routes.[2]
    pkg_path: server/api/routes/[2]
    pkgPath: server/Login/output
    type: type.lib.types.User.views.AuthBasicUserLoginResponse
  parent: server
  parent_path: dsl.server.api
  path: routes
  pkg_path: server/api/routes
  pkgPath: server/Login
  route: login
user-type: type.lib.types.User


*/

// HOFSTADTER_BELOW
