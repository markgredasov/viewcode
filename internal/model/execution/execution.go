package execution

type Language string

type Execution struct {
	Language         Language
	Code             string
	InlineExpression bool
	Arguments        []string
	Environment      map[string]string
}
