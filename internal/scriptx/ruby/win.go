//go:build windows

package ruby

import "github.com/hyprxlabs/run/internal/exec"

func init() {
	exec.Register("ruby", &exec.Executable{
		Name:     "ruby",
		Variable: "RUN_WIN_RUBY_EXE",
		Windows: []string{
			"${ProgramFiles}\\Ruby\\bin\\ruby.exe",
			"${ProgramFiles(x86)}\\Ruby\\bin\\ruby.exe",
		},
	})
}
