package dotnet

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hyprxlabs/run/internal/exec"
	"github.com/hyprxlabs/run/internal/runtimes"
)

const (
	NAME           = "dotnet"
	EVAL_SUPPORTED = false
	executable     = "dotnet"
)

func NewFileCommand(file string, options ...runtimes.RuntimeCmdOption) *runtimes.RuntimeCmd {
	exe, _ := exec.Find(executable, nil)
	if exe == "" {
		exe = executable
	}

	params := &runtimes.RuntimeCmdParams{}
	for _, option := range options {
		option(params)
	}

	args := params.Args

	ext := filepath.Ext(file)
	switch ext {
	case ".dll", ".exe":
		// dotnet <file.dll> [args...]
		splat := []string{file}
		splat = append(splat, args...)
		params.Args = splat
		return runtimes.New(exe, params)

	case ".csproj", ".fsproj", ".vbproj":
		// dotnet run --project <file.csproj> -- [args...]
		splat := []string{"run", "--project", file}
		if len(args) > 0 {
			splat = append(splat, "--")
			splat = append(splat, args...)
		}
		params.Args = splat
		return runtimes.New(exe, params)
	default:
		// dotnet run <file> [args...]
		splat := []string{"run", file}
		splat = append(splat, args...)
		params.Args = splat
		return runtimes.New(exe, params)
	}
}

func NewCodeCommand(code []byte, options ...runtimes.RuntimeCmdOption) *runtimes.RuntimeCmd {
	exe, _ := exec.Find(executable, nil)
	if exe == "" {
		exe = executable
	}

	params := &runtimes.RuntimeCmdParams{}
	for _, option := range options {
		option(params)
	}

	hash := sha256.Sum256(code)

	tempDir := os.TempDir()
	filePath := tempDir + string(os.PathSeparator) + "run-csharp-" + fmt.Sprintf("%x", hash[:8]) + ".cs"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.WriteFile(filePath, code, 0644)
		if err != nil {
			cmd := runtimes.New(exe, params)
			cmd.Err = err
			return cmd
		}
	}

	splat := []string{"run", filePath}
	splat = append(splat, params.Args...)
	params.Args = splat
	cmd := runtimes.New(exe, params)
	cmd.TmpFile = &filePath

	return cmd
}

func NewEvalCommand(code []byte, options ...runtimes.RuntimeCmdOption) *runtimes.RuntimeCmd {
	return NewCodeCommand(code, options...)
}
