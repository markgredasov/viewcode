package handler_executions

import (
	"net/http"

	server_http "github.com/markgredasov/viewcode/internal/server/http"
)

type ExecutionsHTTPHandler struct {
	executionService ExecutionsService
}

type ExecutionsService interface{}

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
