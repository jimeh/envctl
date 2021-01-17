package envctl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWith(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
	}{
		{
			name: "empty",
			env:  map[string]string{},
		},
		{
			name: "new vars",
			env: map[string]string{
				"CC_TEST_INSIDE": "new var is here",
				"CC_TEST_FOO":    "bar",
			},
		},
		{
			name: "set existing",
			env: map[string]string{
				"CC_TEST_INSIDE":  "new var is here",
				"CC_TEST_OUTSIDE": "this is not the same",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("CC_TEST_OUTSIDE", t.Name())
			outside := parseEnviron(os.Environ())

			With(tt.env, func() {
				os.Setenv("CC_TEST_SET_INSIDE", t.Name()+"-manual")

				for k, v := range tt.env {
					got := os.Getenv(k)
					assert.Equal(t, v, got)
				}

				for k, v := range outside {
					if _, ok := tt.env[k]; !ok {
						assert.Equal(t, v, os.Getenv(k))
					}
				}
			})

			_, exists := os.LookupEnv("CC_TEST_SET_INSIDE")
			assert.Equal(t, false, exists)

			for k := range tt.env {
				if _, ok := outside[k]; !ok {
					_, exists := os.LookupEnv(k)
					assert.Equal(t, false, exists)
				}
			}

			for k, v := range outside {
				assert.Equal(t, v, os.Getenv(k))
			}

			os.Unsetenv("CC_TEST_OUTSIDE")
		})
	}
}

func TestWithClean(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
	}{
		{
			name: "empty",
			env:  map[string]string{},
		},
		{
			name: "new vars",
			env: map[string]string{
				"CC_TEST_INSIDE": "new var is here",
				"CC_TEST_FOO":    "bar",
			},
		},
		{
			name: "set existing",
			env: map[string]string{
				"CC_TEST_INSIDE":  "new var is here",
				"CC_TEST_OUTSIDE": "is is not the same",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("CC_TEST_OUTSIDE", t.Name())
			outside := parseEnviron(os.Environ())

			WithClean(tt.env, func() {
				os.Setenv("CC_TEST_SET_INSIDE", t.Name()+"-manual")

				for k, v := range tt.env {
					got := os.Getenv(k)
					assert.Equal(t, v, got)
				}

				for k := range outside {
					if _, ok := tt.env[k]; !ok {
						_, exists := os.LookupEnv(k)
						assert.Equal(t, false, exists)
					}
				}
			})

			_, exists := os.LookupEnv("CC_TEST_SET_INSIDE")
			assert.Equal(t, false, exists)

			for k := range tt.env {
				if _, ok := outside[k]; !ok {
					_, exists := os.LookupEnv(k)
					assert.Equal(t, false, exists)
				}
			}

			for k, v := range outside {
				assert.Equal(t, v, os.Getenv(k))
			}

			os.Unsetenv("CC_TEST_OUTSIDE")
		})
	}
}

func Test_parseEnviron(t *testing.T) {
	tests := []struct {
		name string
		vars []string
		want map[string]string
	}{
		{
			name: "empty",
			vars: []string{},
			want: map[string]string{},
		},
		{
			name: "various",
			vars: []string{
				"USER=john",
				"NAME=John Doe",
				"SHELL=/bin/bash",
				"GOPRIVATE=",
				"LESS= -R",
				"LESSOPEN=| src-hilite-lesspipe.sh %s",
				"X=11",
				"TAGS=foo=bar,hello=world",
				"_=go",
				"INVALID-var",
			},
			want: map[string]string{
				"USER":      "john",
				"NAME":      "John Doe",
				"SHELL":     "/bin/bash",
				"GOPRIVATE": "",
				"LESS":      " -R",
				"LESSOPEN":  "| src-hilite-lesspipe.sh %s",
				"X":         "11",
				"TAGS":      "foo=bar,hello=world",
				"_":         "go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseEnviron(tt.vars)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Benchmark_parseEnviron(b *testing.B) {
	env := []string{
		`EDITOR=emacsclient-wrapper`,
		`EMACS=/Applications/Emacs.app/Contents/MacOS/Emacs`,
		`GEM_EDITOR=emacsclient-wrapper`,
		`GOPATH=/Users/jimeh/.go`,
		`HOME=/Users/jimeh`,
		`HOMEBREW_NO_ANALYTICS=1`,
		`LANG=en_US.UTF-8`,
		`LC_ALL=en_US.UTF-8`,
		`LC_TERMINAL=iTerm2`,
		`LC_TERMINAL_VERSION=3.4.3`,
		`LESS= -R`,
		`LESSOPEN=| src-hilite-lesspipe.sh %s`,
		`NODENV_SHELL=zsh`,
		`PWD=/Users/jimeh/Projects/envctl`,
		`PYENV_SHELL=zsh`,
		`RBENV_SHELL=zsh`,
		`TERM=screen-256color`,
		`TERM_PROGRAM=iTerm.app`,
		`TERM_PROGRAM_VERSION=3.4.3`,
		`TMPDIR=/tmp/user-jimeh`,
		`TMUX=/private/tmp/tmux-501/default,4148,2`,
		`TMUX_PANE=%29`,
		`TMUX_PLUGIN_MANAGER_PATH=/Users/jimeh/.tmux/plugins/`,
		`USER=jimeh`,
		`ZPFX=/Users/jimeh/.local/zsh/zinit/polaris`,
		`ZSH_CACHE_DIR=/Users/jimeh/.cache/zinit`,
		`_=/usr/bin/env`,
		`__CFBundleIdentifier=com.googlecode.iterm2`,
	}

	for i := 0; i < b.N; i++ {
		parseEnviron(env)
	}
}
