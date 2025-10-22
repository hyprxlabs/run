//go:build windows

package golang

import "github.com/hyprxlabs/run/internal/exec"

func init() {
	exec.Register(executable, &exec.Executable{
		Name:     executable,
		Variable: "RUN_WIN_GOLANG_EXE",
		Windows: []string{
			"${LOCALAPPDATA}\\Go\\bin\\go.exe",
			"${ProgramFiles}\\Go\\bin\\go.exe",
			"${ProgramFiles(x86)}\\Go\\bin\\go.exe",
		},
	})
}
