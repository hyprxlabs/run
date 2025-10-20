package schema

import "go.yaml.in/yaml/v4"

type Hooks struct {
	Before []string
	After  []string
}

func (h *Hooks) init() {
	if h == nil {
		h = &Hooks{
			Before: []string{},
			After:  []string{},
		}
		return
	}

	if h.Before == nil {
		h.Before = []string{}
	}

	if h.After == nil {
		h.After = []string{}
	}
}

func (h *Hooks) UnmarshalYAML(node *yaml.Node) error {
	h.init()

	if node.Kind == yaml.ScalarNode {
		if node.Value == "true" || node.Value == "yes" {
			h.Before = []string{"before"}
			h.After = []string{"after"}
			return nil
		}

		if node.Value == "false" || node.Value == "no" {
			h.Before = []string{}
			h.After = []string{}
			return nil
		}

		return yamlErrorf(*node, "expected 'true' or 'false' for hooks scalar")
	}

	if node.Kind != yaml.MappingNode {
		return yamlErrorf(*node, "expected yaml mapping for hooks")
	}

	for i := 0; i < len(node.Content); i += 2 {
		keyNode := node.Content[i]
		valNode := node.Content[i+1]

		switch keyNode.Value {
		case "before":
			switch valNode.Kind {
			case yaml.ScalarNode:
				h.Before = []string{valNode.Value}
			case yaml.SequenceNode:
				var before []string
				err := valNode.Decode(&before)
				if err != nil {
					return yamlErrorf(*valNode, "failed to decode 'before' hooks: %v", err)
				}
				h.Before = before
			default:
				return yamlErrorf(*valNode, "expected yaml scalar or sequence for 'before' hooks")
			}
		case "after":
			switch valNode.Kind {
			case yaml.ScalarNode:
				h.After = []string{valNode.Value}
			case yaml.SequenceNode:
				var after []string
				err := valNode.Decode(&after)
				if err != nil {
					return yamlErrorf(*valNode, "failed to decode 'after' hooks: %v", err)
				}
				h.After = after
			default:
				return yamlErrorf(*valNode, "expected yaml scalar or sequence for 'after' hooks")
			}
		default:
			// Ignore unknown fields for forward compatibility
		}
	}

	return nil
}
