package post

import (
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"

	"github.com/hofstadter-io/examples/blog/server/resources/user/post/comment"
)

/* no when
Name: post
Parent: user
pkgPath: server/user/post
*/

var logger = log.New()

func SetLogger(l log.Logger) {
	ldcfg := viper.GetStringMap("log-config.resources.user.post.default")

	if ldcfg == nil || len(ldcfg) == 0 {
		logger = l
	} else {
		// find the logging level
		level_str := ldcfg["level"].(string)
		level, err := log.LvlFromString(level_str)
		if err != nil {
			panic(err)
		}

		// possibly find the stack switch
		stack := false
		stack_tmp := ldcfg["stack"]
		if stack_tmp != nil {
			stack = stack_tmp.(bool)
		}

		// build the local logger
		termlog := log.LvlFilterHandler(level, log.StdoutHandler)
		if stack {
			term_stack := log.CallerStackHandler("%+v", log.StdoutHandler)
			termlog = log.LvlFilterHandler(level, term_stack)
		}

		// set the local logger
		logger.SetHandler(termlog)
	}

	// set sub-command loggers before possibly overriding locally next
	setSubLoggers(logger)

	// possibly override locally
	lcfg := viper.GetStringMap("log-config.resources.user.post")

	if lcfg == nil || len(lcfg) == 0 {
		logger = l
	} else {
		// find the logging level
		level_str := lcfg["level"].(string)
		level, err := log.LvlFromString(level_str)
		if err != nil {
			panic(err)
		}

		// possibly find the stack switch
		stack := false
		stack_tmp := lcfg["stack"]
		if stack_tmp != nil {
			stack = stack_tmp.(bool)
		}

		// build the local logger
		termlog := log.LvlFilterHandler(level, log.StdoutHandler)
		if stack {
			term_stack := log.CallerStackHandler("%+v", log.StdoutHandler)
			termlog = log.LvlFilterHandler(level, term_stack)
		}

		// set the local logger
		logger.SetHandler(termlog)
	}

}

func setSubLoggers(logger log.Logger) {

	// ----
	comment.SetLogger(logger)

}

// HOFSTADTER_BELOW
