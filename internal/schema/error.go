package schema

import (
	"fmt"

	"github.com/hyprxlabs/run/internal/errors"
	"go.yaml.in/yaml/v4"
)

func yamlError(node yaml.Node, message string) error {
	msg := fmt.Sprintf("%s on line %d, at column %d", message, node.Line, node.Column)

	return errors.NewDetails(msg, "YamlError", msg)
}

func yamlErrorf(node yaml.Node, format string, args ...interface{}) error {
	return yamlError(node, fmt.Sprintf(format, args...))
}
