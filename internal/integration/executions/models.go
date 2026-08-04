package integration_executions

type executionsRequest struct {
	Language         string            `json:"language"`
	Code             string            `json:"code"`
	InlineExpression bool              `json:"inline_expression"`
	Arguments        []string          `json:"arguments"`
	Environment      map[string]string `json:"environment"`
}

type executionsResponse struct {
	StdOut     string `json:"stdout"`
	StdErr     string `json:"stderr"`
	ExecTimeMS int    `json:"exectimems"`
}
