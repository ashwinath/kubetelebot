package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunShell(t *testing.T) {
	var tests = []struct {
		name     string
		binary   string
		args     []string
		expected string
	}{
		{
			"echo hello world",
			"echo",
			[]string{"-n", "hello world"},
			"hello world",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			res, err := RunShell(tt.binary, tt.args...)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, *res)
		})
	}
}
