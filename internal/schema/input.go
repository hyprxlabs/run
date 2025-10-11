package schema

type Input struct {
	Id        string
	Name      *string
	Desc      *string
	Type      *string
	Default   interface{}
	Required  *bool
	Selection []string
}
