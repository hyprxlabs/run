package schema

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Runfile struct {
	Name string

	Config      RunfileConfig
	Path        string
	Env         Environment
	DotEnv      []string
	Tasks       Tasks
	HostImports HostImports
	Values      map[string]interface{}
	Args        []string
}

func NewRunfile() *Runfile {
	defaultContext := os.Getenv("RUN_CONTEXT")
	if len(defaultContext) == 0 {
		defaultContext = "default"
	}

	defaultShell := os.Getenv("RUN_SHELL")
	if len(defaultShell) == 0 {
		defaultShell = "bash"
		if os.Getenv("OS") == "Windows_NT" {
			defaultShell = "powershell"
		}
	}

	return &Runfile{
		Name: "",
		Config: RunfileConfig{
			Substitution: true,
			Dirs: Dirs{
				Etc:      "./.run/etc",
				Projects: []string{"./.run/apps"},
			},
			Paths:   Paths{},
			Env:     *NewEnv(),
			Shell:   &defaultShell,
			Context: &defaultContext,
		},
		Env:         *NewEnv(),
		DotEnv:      []string{},
		Tasks:       *NewTasks(),
		HostImports: *NewHostImports(),
		Values:      map[string]interface{}{},
	}
}

func (x *Runfile) DecodeYAMLFile(path string) error {
	if x == nil {
		x = NewRunfile()
	}

	if !filepath.IsAbs(path) {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		path = absPath
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, x); err != nil {
		return err
	}

	x.Path = path
	return nil
}
