package ruby

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"

	"github.com/hyprxlabs/run/internal/exec"
	"github.com/hyprxlabs/run/internal/runtimes"
)

const (
	NAME           = "ruby"
	EVAL_SUPPORTED = true
	executable     = "ruby"
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

func NewScriptCommand(script string, options ...runtimes.RuntimeCmdOption) *runtimes.RuntimeCmd {
	exe, _ := exec.Find(executable, nil)
	if exe == "" {
		exe = executable
	}

	if !strings.ContainsAny(script, "\n\r") {
		trimmed := strings.TrimSpace(script)
		if strings.HasSuffix(trimmed, ".rb") {
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

	splat := []string{file}
	params := &runtimes.RuntimeCmdParams{}
	for _, option := range options {
		option(params)
	}

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
	filePath := tempDir + string(os.PathSeparator) + "run-ruby-" + fmt.Sprintf("%x", hash[:8]) + ".rb"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.WriteFile(filePath, code, 0644)
		if err != nil {
			cmd := runtimes.NewWithOptions(exe, options...)
			cmd.Err = err
			return cmd
		}
	}

	splat := []string{filePath}
	params := &runtimes.RuntimeCmdParams{}
	for _, option := range options {
		option(params)
	}

	splat = append(splat, params.Args...)
	params.Args = splat
	cmd := runtimes.New(exe, params)
	cmd.TmpFile = &filePath
	return cmd
}

func EvalCodeCommand(code []byte, options ...runtimes.RuntimeCmdOption) *runtimes.RuntimeCmd {
	exe, _ := exec.Find("ruby", nil)
	if exe == "" {
		exe = "ruby"
	}

	params := &runtimes.RuntimeCmdParams{}
	for _, option := range options {
		option(params)
	}

	args := params.Args

	splat := []string{"-e", string(code)}
	splat = append(splat, args...)
	params.Args = splat
	return runtimes.New(exe, params)
}
