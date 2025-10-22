package runtimes

import (
	"context"
	"os"

	"github.com/hyprxlabs/run/internal/exec"
)

type RuntimeCmd struct {
	*exec.Cmd
	TmpFile *string
}

func (r *RuntimeCmd) Cleanup() {
	if r.TmpFile != nil {
		os.Remove(*r.TmpFile)
	}
}

type RuntimeCmdParams struct {
	Args    []string
	TmpFile *string
	Cwd     *string
	Env     map[string]string
	Context context.Context
}

type RuntimeCmdOption func(*RuntimeCmdParams)

func WithArgs(args ...string) RuntimeCmdOption {
	return func(p *RuntimeCmdParams) {
		p.Args = args
	}
}

func WithTmpFile(tmpFile string) RuntimeCmdOption {
	return func(p *RuntimeCmdParams) {
		p.TmpFile = &tmpFile
	}
}

func WithCwd(cwd string) RuntimeCmdOption {
	return func(p *RuntimeCmdParams) {
		p.Cwd = &cwd
	}
}

func WithEnv(env map[string]string) RuntimeCmdOption {
	return func(p *RuntimeCmdParams) {
		p.Env = env
	}
}

func WithContext(ctx context.Context) RuntimeCmdOption {
	return func(p *RuntimeCmdParams) {
		p.Context = ctx
	}
}

func New(exe string, params *RuntimeCmdParams) *RuntimeCmd {
	if params == nil {
		params = &RuntimeCmdParams{}
	}

	var rc *RuntimeCmd
	if params.Context != nil {
		rc = &RuntimeCmd{
			Cmd: exec.NewContext(params.Context, exe, params.Args...),
		}
	} else {
		rc = &RuntimeCmd{
			Cmd: exec.New(exe, params.Args...),
		}
	}

	if params.TmpFile != nil {
		rc.TmpFile = params.TmpFile
	}
	if params.Cwd != nil {
		rc.Cmd.WithCwd(*params.Cwd)
	}
	if params.Env != nil {
		rc.Cmd.WithEnvMap(params.Env)
	}
	return rc
}

func NewWithOptions(exe string, options ...RuntimeCmdOption) *RuntimeCmd {
	params := &RuntimeCmdParams{}
	for _, option := range options {
		option(params)
	}
	return New(exe, params)
}
