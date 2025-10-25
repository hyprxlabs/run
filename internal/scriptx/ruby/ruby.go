package ruby

const NAME = "ruby"

var ScriptArgs = []string{"-e"}

func New(args ...string) *exec.Cmd {
	return exec.New(NAME, args...)
}

func NewContext(ctx context.Context, args ...string) *exec.Cmd {
	return exec.NewContext(ctx, NAME, args...)
}

func File(file string, args ...string) *exec.Cmd {
	allArgs := append([]string{file}, args...)
	return exec.New(NAME, allArgs...)
}

func FileContext(ctx context.Context, file string, args ...string) *exec.Cmd {
	allArgs := append([]string{file}, args...)
	return exec.NewContext(ctx, NAME, allArgs...)
}

func Inline(script string, args ...string) *exec.Cmd {
	allArgs := append(append(ScriptArgs, script), args...)
	return exec.New(NAME, allArgs...)
}

func InlineContext(ctx context.Context, script string, args ...string) *exec.Cmd {
	allArgs := append(append(ScriptArgs, script), args...)
	return exec.NewContext(ctx, NAME, allArgs...)
}

func Script(script string, args ...string) *exec.Cmd {
	if !strings.ContainsAny(script, "\n\r") {
		trimmed := strings.TrimSpace(script)
		if strings.HasSuffix(trimmed, ".rb") {
			return File(trimmed, args...)
		}
	}

	return Inline(script, args...)
}

func ScriptContext(ctx context.Context, script string, args ...string) *exec.Cmd {
	if !strings.ContainsAny(script, "\n\r") {
		trimmed := strings.TrimSpace(script)
		if strings.HasSuffix(trimmed, ".rb") {
			return FileContext(ctx, trimmed, args...)
		}
	}

	return InlineContext(ctx, script, args...)
}