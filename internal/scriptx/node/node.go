package node 

const NAME = "node"

var Extensions = []string{".js", ".mjs", ".cjs", ".ts"}

var ScriptArgs = []string{}

func New(args ...string) *exec.Cmd {
	return exec.New(NAME, args...)
}

func NewContext(ctx context.Context, args ...string) *exec.Cmd {
	return exec.NewContext(ctx, NAME, args...)
}

func File(path string, args ...string) *exec.Cmd {
	allArgs := append(append(ScriptArgs, path), args...)
	return exec.New(NAME, allArgs...)
}

func FileContext(ctx context.Context, path string, args ...string) *exec.Cmd {
	allArgs := append(append(ScriptArgs, path), args...)
	return exec.NewContext(ctx, NAME, allArgs...)
}

func Inline(script string, args ...string) *exec.Cmd {
	splat := append(ScriptArgs, "-e", script)
	allArgs := append(splat, args...)
	return exec.New(NAME, allArgs...)
}

func InlineContext(ctx context.Context, script string, args ...string) *exec.Cmd {
	splat := append(ScriptArgs, "-e", script)
	allArgs := append(splat, args...)
	return exec.NewContext(ctx, NAME, allArgs...)
}

func Script(script string, args ...string) *exec.Cmd {
	if !strings.ContainsAny(script, "\n\r") {
		trimmed := strings.TrimSpace(script)
		for _, ext := range Extensions {
			if strings.HasSuffix(trimmed, ext) {
				return File(trimmed, args...)
			}
		}
	}

	return Inline(script, args...)
}

func ScriptContext(ctx context.Context, script string, args ...string) *exec.Cmd {
	if !strings.ContainsAny(script, "\n\r") {
		trimmed := strings.TrimSpace(script)
		for _, ext := range Extensions {
			if strings.HasSuffix(trimmed, ext) {
				return FileContext(ctx, trimmed, args...)
			}
		}
	}

	return InlineContext(ctx, script, args...)
}