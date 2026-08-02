package app

import (
	handler_executions "github.com/markgredasov/viewcode/internal/handler/http/executions"
	server_http "github.com/markgredasov/viewcode/internal/server/http"
)

func (a *App) initFeatures() ([]server_http.Route, error) {
	routes := a.initExecutionFeature()
	return routes, nil
}

func (a *App) initExecutionFeature() []server_http.Route {
	executionHandler := handler_executions.NewExecutionHTTPHandler(nil)
	return executionHandler.Routes()
}
