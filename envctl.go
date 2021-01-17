// Package envctl provides test helper functions that temporarily changes
// environment variables accessible via os.Getenv() and os.LookupEnv(). After
// they are done, all changes to environment variables are reset, including
// manual changes with os.Setenv() within the helper functions.
package envctl

import (
	"os"
	"strings"
	"sync"
)

var mux = &sync.Mutex{}

// With temporarily sets all given vars as environment variables during the
// execution of f function. Existing environment variables are also available
// within f. Any overridden environment variables will contain the overridden
// value..
//
// After f execution completes all changes to environment variables are reset,
// including manual changes within the f function.
func With(env map[string]string, f func()) {
	mux.Lock()
	defer mux.Unlock()

	undo := parseEnviron(os.Environ())

	apply(env)
	defer func() {
		os.Clearenv()
		apply(undo)
	}()

	f()
}

// WithClean temporarily changes all environment variables available within f
// function to only be those provided. Existing environment variables are not
// available within f.
//
// After f execution completes all changes to environment variables are reset,
// including manual changes within the f function.
func WithClean(env map[string]string, f func()) {
	mux.Lock()
	defer mux.Unlock()

	undo := parseEnviron(os.Environ())

	os.Clearenv()
	apply(env)
	defer func() {
		os.Clearenv()
		apply(undo)
	}()

	f()
}

func apply(vars map[string]string) {
	for k, v := range vars {
		os.Setenv(k, v)
	}
}

func parseEnviron(vars []string) map[string]string {
	r := map[string]string{}

	for _, v := range vars {
		i := strings.Index(v, "=")
		if i < 1 {
			continue
		}

		r[v[0:i]] = v[i+1:]
	}

	return r
}
