package climatecontrol

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithEnv(t *testing.T) {
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

			WithEnv(tt.env, func() {
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

func TestWithCleanEnv(t *testing.T) {
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

			WithCleanEnv(tt.env, func() {
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
