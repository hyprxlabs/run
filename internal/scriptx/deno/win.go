//go:build windows

package deno

import "github.com/hyprxlabs/run/internal/exec"

func init() {
	exec.Register("deno", &exec.Executable{
		Name:     "deno",
		Variable: "RUN_DENO_EXE",
		Windows: []string{
			"C:\\Program Files\\deno\\deno.exe",
			"C:\\deno\\deno.exe",
		},
	})
}