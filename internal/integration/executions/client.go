package integration_executions

import (
	"context"
	"fmt"
	"time"

	ranna "github.com/ranna-go/ranna/pkg/client"
	"github.com/ranna-go/ranna/pkg/models"
)

const (
	execURL = "http://localhost:8081"
	timeout = "10s"
)

type client struct {
	rannaClient ranna.Client
	execTimeout time.Duration
}

func newClient() (*client, error) {
	rannaClient, err := ranna.New(ranna.Options{
		Endpoint: execURL,
	})

	if err != nil {
		return nil, err
	}

	timeout, err := time.ParseDuration(timeout)
	if err != nil {
		return nil, err
	}

	return &client{
		rannaClient: rannaClient,
		execTimeout: timeout,
	}, nil
}

func (c *client) makeExecution(ctx context.Context, language, code string) (*executionsResponse, error) {
	reqBody := executionsRequest{
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

	return &executionsResponse{
		StdOut:     resp.StdOut,
		StdErr:     resp.StdErr,
		ExecTimeMS: resp.ExecTimeMS,
	}, nil
}
