package workflows

import "github.com/hyprxlabs/run/internal/schema"

type CyclicalReferenceError struct {
	Cycles []schema.Task
}

func (e *CyclicalReferenceError) Error() string {
	msg := "Cyclical references found in tasks:\n"
	for _, cycle := range e.Cycles {
		msg += " - " + cycle.Id + "\n"
	}
	return msg
}
