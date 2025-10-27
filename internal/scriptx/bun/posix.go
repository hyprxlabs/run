//go:build !windows

package bun

import "github.com/hyprxlabs/run/internal/exec"

func init() {
	exec.Register("bun", &exec.Executable{
		Name:     "bun",
		Variable: "RUN_BUN_EXE",
		Linux: []string{
			"${HOME}/.bun/bin/bun",
			"${HOME}/.local/bin/bun",
			"/usr/bin/bun",
			"/usr/local/bin/bun",
		},
	})
}
