//go:build !windows

package bash

import "github.com/hyprxlabs/run/internal/exec"

func init() {
	exec.Register("bash", &exec.Executable{
		Name:     "bash",
		Variable: "RUN_BASH_EXE",
		Linux: []string{
			"/bin/bash",
			"/usr/bin/bash",
		},
	})
}


func resolveScriptFile(script string) string {
	return script 
}