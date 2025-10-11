package tasks

import (
	"context"
	"time"

	"github.com/hyprxlabs/run/internal/schema"
)

type TaskModel struct {
	Id      string
	Name    string
	Help    string
	Env     schema.Environment
	Desc    string
	Run     string
	Uses    string
	Args    []string
	Hosts   schema.Hosts
	Cwd     string
	Timeout time.Duration
	Needs   []string
	With    schema.With
}

type TaskContext struct {
	Schema      *schema.Task
	Task        *TaskModel
	Context     context.Context
	Args        []string
	ContextName string
}

type TaskHandler func(tc TaskContext) *TaskResult

type TaskHandlerRegistry map[string]TaskHandler

var GlobalTaskHandlers = TaskHandlerRegistry{
	"ssh":        runSshTask,
	"tmpl":       runTpl,
	"scp":        runSCP,
	"bash":       runShell,
	"sh":         runShell,
	"zsh":        runShell,
	"powershell": runShell,
	"pwsh":       runShell,
	"cmd":        runShell,
	"python":     runShell,
	"ruby":       runShell,
	"deno":       runShell,
	"node":       runShell,
	"bun":        runShell,
	"shell":      runShell,
}
