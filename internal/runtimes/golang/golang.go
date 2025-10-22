package golang

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"

	"github.com/hyprxlabs/run/internal/exec"
	"github.com/hyprxlabs/run/internal/runtimes"
)

const (
	NAME           = "go"
	EVAL_SUPPORTED = false
	executable     = "go"
)

func New(args ...string) *runtimes.RuntimeCmd {
	exe, _ := exec.Find(executable, nil)
	if exe == "" {
		exe = executable
	}

	if len(args) == 0 {
		return runtimes.NewWithOptions(exe)
	}

	return runtimes.NewWithOptions(exe, runtimes.WithArgs(args...))
}

// NewScriptCommand can handle both a script file path or raw Go code.
// the script file must not include any new lines and must end with .go to be treated as a file path.
func NewScriptCommand(script string, options ...runtimes.RuntimeCmdOption) *runtimes.RuntimeCmd {
	exe, _ := exec.Find(executable, nil)
	if exe == "" {
		exe = executable
	}

	if !strings.ContainsAny(script, "\n\r") {
		trimmed := strings.TrimSpace(script)
		if strings.HasSuffix(trimmed, ".go") {
			return NewFileCommand(trimmed, options...)
		}
	}

	return NewCodeCommand([]byte(script), options...)
}

func NewFileCommand(file string, options ...runtimes.RuntimeCmdOption) *runtimes.RuntimeCmd {
	exe, _ := exec.Find(executable, nil)
	if exe == "" {
		exe = executable
	}

	params := &runtimes.RuntimeCmdParams{}
	for _, option := range options {
		option(params)
	}
	splat := []string{"run", file}
	splat = append(splat, params.Args...)
	params.Args = splat
	return runtimes.New(exe, params)
}

func NewCodeCommand(code []byte, options ...runtimes.RuntimeCmdOption) *runtimes.RuntimeCmd {
	exe, _ := exec.Find(executable, nil)
	if exe == "" {
		exe = executable
	}

	hash := sha256.Sum256(code)

	tempDir := os.TempDir()
	filePath := tempDir + string(os.PathSeparator) + "run-go-" + fmt.Sprintf("%x", hash[:8]) + ".go"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.WriteFile(filePath, code, 0644)
		if err != nil {
			cmd := runtimes.New(exe, &runtimes.RuntimeCmdParams{})
			cmd.Err = err
			return cmd
		}
	}

	params := &runtimes.RuntimeCmdParams{}
	for _, option := range options {
		option(params)
	}
	splat := []string{"run", filePath}
	splat = append(splat, params.Args...)
	params.Args = splat
	cmd := runtimes.New(exe, params)
	return cmd
}

func NewEvalCommand(code []byte, options ...runtimes.RuntimeCmdOption) *runtimes.RuntimeCmd {
	return NewCodeCommand(code, options...)
}
