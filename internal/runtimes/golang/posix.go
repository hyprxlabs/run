//go:build !windows

package golang

import "github.com/hyprxlabs/run/internal/exec"

func init() {
	exec.Register("go", &exec.Executable{
		Name:     "go",
		Variable: "RUN_GO_EXE",
		Linux: []string{
			"${HOME}/.go/bin/go",
			"/usr/local/go/bin/go",
			"/usr/bin/go",
			"/usr/local/bin/go",
		},
	})
}
