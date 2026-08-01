package execution

import (
	"context"
	"fmt"

	ranna "github.com/ranna-go/ranna/pkg/client"
	"github.com/ranna-go/ranna/pkg/models"
)

const (
	execURL = "/"
)

type client struct {
	rannaClient ranna.Client
}

func newClient() (*client, error) {
	rannaClient, err := ranna.New(ranna.Options{
		Endpoint: execURL,
	})

	if err != nil {
		return nil, err
	}

	return &client{
		rannaClient: rannaClient,
	}, nil
}

func (c *client) makeExecution(ctx context.Context, language, code string) (*executionResponse, error) {
	reqBody := executionRequest{
		Language:         language,
		Code:             code,
		InlineExpression: false,
		Arguments:        []string{},
		Environment:      make(map[string]string),
	}

	resp, err := c.rannaClient.Exec(models.ExecutionRequest{
		Language:         reqBody.Language,
		Code:             reqBody.Code,
		InlineExpression: reqBody.InlineExpression,
		Arguments:        reqBody.Arguments,
		Environment:      reqBody.Environment,
	})

	if err != nil {
		return nil, fmt.Errorf("exec: %w", err)
	}

	return &executionResponse{
		StdOut:     resp.StdOut,
		StdErr:     resp.StdErr,
		ExecTimeMS: resp.ExecTimeMS,
	}, nil
}
