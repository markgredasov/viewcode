package handler_executions

import (
	"context"
	"net/http"

	model_execution "github.com/markgredasov/viewcode/internal/model/execution"
	server_http "github.com/markgredasov/viewcode/internal/server/http"
)

type ExecutionsHTTPHandler struct {
	executionService ExecutionsService
}

type ExecutionsService interface {
	MakeExecution(ctx context.Context, language, code string) (model_execution.Result, error)
}

func NewExecutionHTTPHandler(service ExecutionsService) *ExecutionsHTTPHandler {
	return &ExecutionsHTTPHandler{
		executionService: service,
	}
}

func (h *ExecutionsHTTPHandler) Routes() []server_http.Route {
	return []server_http.Route{
		{
			Method:  http.MethodPost,
			Path:    "/executions",
			Handler: h.MakeExecution,
		},
	}
}
