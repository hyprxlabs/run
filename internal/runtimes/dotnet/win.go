//go:build windows

package dotnet

import "github.com/hyprxlabs/run/internal/exec"

func init() {
	exec.Register(executable, &exec.Executable{
		Name:     executable,
		Variable: "RUN_WIN_DOTNET_EXE",
		Windows: []string{
			"${LOCALAPPDATA}\\Microsoft\\dotnet\\dotnet.exe",
			"${ProgramFiles}\\dotnet\\dotnet.exe",
			"${ProgramFiles(x86)}\\dotnet\\dotnet.exe",
		},
	})
}
