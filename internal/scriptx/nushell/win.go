//go:build windows

package nushell

import (
	"github.com/hyprxlabs/run/internal/exec"
)

func init() {
	exec.Register("nu", &exec.Executable{
		Name:     "nu",
		Variable: "RUN_WIN_NUSHELL_EXE",
		Windows: []string{
			"C:\\Program Files\\nu\\bin\\nu.exe",
		},
	})
}
