//go:build !windows

package dotnet

import "github.com/hyprxlabs/run/internal/exec"

func init() {
	exec.Register(executable, &exec.Executable{
		Name:     executable,
		Variable: "RUN_DOTNET_EXE",
		Linux: []string{
			"${HOME}/.dotnet/dotnet",
			"/usr/bin/dotnet",
			"/usr/local/bin/dotnet",
		},
	})
}
