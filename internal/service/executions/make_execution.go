package service_executions

import (
	"context"
	"fmt"

	model_execution "github.com/markgredasov/viewcode/internal/model/execution"
)

func (es *ExecutionsService) MakeExecution(ctx context.Context, language, code string) (model_execution.Result, error) {
	executionEntity, err := model_execution.NewExecution(language, code)
	if err != nil {
		return model_execution.Result{}, fmt.Errorf("create execution entity: %w", err)
	}

	return es.execPort.MakeExecution(ctx, executionEntity)
}
