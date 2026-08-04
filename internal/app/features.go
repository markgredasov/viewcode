package app

import (
	handler_executions "github.com/markgredasov/viewcode/internal/handler/http/executions"
	integration_executions "github.com/markgredasov/viewcode/internal/integration/executions"
	server_http "github.com/markgredasov/viewcode/internal/server/http"
	service_executions "github.com/markgredasov/viewcode/internal/service/executions"
)

func (a *App) initFeatures() ([]server_http.Route, error) {
	return a.initExecutionFeature()
}

func (a *App) initExecutionFeature() ([]server_http.Route, error) {
	executionsAdapter, err := integration_executions.NewAdapter()
	if err != nil {
		return []server_http.Route{}, err
	}
	executionsService := service_executions.NewExecutionsService(executionsAdapter)
	executionsHandler := handler_executions.NewExecutionHTTPHandler(&executionsService)
	return executionsHandler.Routes(), nil
}
