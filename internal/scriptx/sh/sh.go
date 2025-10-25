package sh

const NAME = "sh"

var ScriptArgs = []string{"-e"}

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
	splat := append(ScriptArgs, script)
	splat = append(splat, args...)
	return exec.New(NAME, splat...)
}

func InlineContext(ctx context.Context, script string, args ...string) *exec.Cmd {
	splat := append(ScriptArgs, script)
	splat = append(splat, args...)
	return exec.NewContext(ctx, NAME, splat...)
}

func Script(script string, args ...string) *exec.Cmd {
	if !strings.ContainsAny(script, "\n\r") {
		trimmed := strings.TrimSpace(script)
		if strings.HasSuffix(trimmed, ".sh") {
			return File(trimmed, args...)
		}
	}

	return Inline(script, args...)
}

func ScriptContext(ctx context.Context, script string, args ...string) *exec.Cmd {
	if !strings.ContainsAny(script, "\n\r") {
		trimmed := strings.TrimSpace(script)
		if strings.HasSuffix(trimmed, ".sh") {
			return FileContext(ctx, trimmed, args...)
		}
	}

	return InlineContext(ctx, script, args...)
}