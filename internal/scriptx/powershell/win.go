//go:build windows

package powershell

import "github.com/hyprxlabs/run/internal/exec"

func init() {
	exec.Register("powershell", &exec.Executable{
		Name:     "powershell",
		Variable: "RUN_WIN_POWERSHELL_EXE",
		Windows: []string{
			"${SystemRoot}\\System32\\WindowsPowerShell\\v1.0\\powershell.exe",
		},
	})
}
