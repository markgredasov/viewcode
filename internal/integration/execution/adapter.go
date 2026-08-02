package execution

import (
	"context"
	"fmt"

	errs "github.com/markgredasov/viewcode/internal/errors"
	model_execution "github.com/markgredasov/viewcode/internal/model/execution"
)

var languageToExecutionLanguage = map[model_execution.Language]string{
	model_execution.LanguageGo:     "go",
	model_execution.LanguagePython: "python3",
	model_execution.LanguageRust:   "rust",
}

type Adapter struct {
	client *client
}

func NewAdapter() (*Adapter, error) {
	c, err := newClient()
	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	return &Adapter{
		client: c,
	}, nil
}

func (a *Adapter) MakeExecution(ctx context.Context, exec model_execution.Execution) (model_execution.Result, error) {
	language, ok := languageToExecutionLanguage[exec.Language]
	if !ok {
		return model_execution.Result{}, fmt.Errorf("convert language: %w", errs.ErrInvalidLanguage)
	}

	execResponse, err := a.client.makeExecution(ctx, language, exec.Code)
	if err != nil {
		return model_execution.Result{}, fmt.Errorf("make exec: %w", err)
	}

	return model_execution.Result{
		StdOut:            execResponse.StdOut,
		StdErr:            execResponse.StdErr,
		ExecutionTimeInMS: execResponse.ExecTimeMS,
	}, nil
}
