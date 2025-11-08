//go:build !windows

package powershell

import "github.com/hyprxlabs/run/internal/exec"

func init() {
	exec.Register("powershell", &exec.Executable{
		Name:     "powershell",
		Variable: "RUN_POWERSHELL_EXE",
		Linux: []string{
			"/bin/powershell",
			"/usr/bin/powershell",
		},
	})
}
