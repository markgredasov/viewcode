package service_executions

import (
	"context"

	model_execution "github.com/markgredasov/viewcode/internal/model/execution"
)

type ExecutionsService struct {
	execPort ExecutionsPort
}

type ExecutionsPort interface {
	MakeExecution(ctx context.Context, exec model_execution.Execution) (model_execution.Result, error)
}

func NewExecutionsService(executionsPort ExecutionsPort) ExecutionsService {
	return ExecutionsService{
		execPort: executionsPort,
	}
}
