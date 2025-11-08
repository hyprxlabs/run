//go:build windows

package python

import "github.com/hyprxlabs/run/internal/exec"

func init() {
	exec.Register("python", &exec.Executable{
		Name:     "python",
		Variable: "RUN_WIN_PYTHON_EXE",
		Windows: []string{
			"${ProgramFiles}\\Python\\Python.exe",
			"${ProgramFiles(x86)}\\Python\\Python.exe",
		},
	})
}
