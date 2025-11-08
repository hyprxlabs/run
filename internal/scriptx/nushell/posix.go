//go:build !windows

package nushell

import "github.com/hyprxlabs/run/internal/exec"

func init() {
	exec.Register("nu", &exec.Executable{
		Name:     "nu",
		Variable: "RUN_NUSHELL_EXE",
		Linux: []string{
			"/bin/nu",
			"/usr/bin/nu",
		},
	})
}
