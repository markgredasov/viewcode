package model_execution

type Execution struct {
	Language         Language
	Code             string
	InlineExpression bool
	Arguments        []string
	Environment      map[string]string
}

type Result struct {
	StdOut            string
	StdErr            string
	ExecutionTimeInMS int
}
