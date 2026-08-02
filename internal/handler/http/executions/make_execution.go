package handler_executions

import (
	"net/http"

	"github.com/markgredasov/viewcode/internal/logger"
	"github.com/markgredasov/viewcode/internal/server/http/response"
)

func (h *ExecutionsHTTPHandler) MakeExecution(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, w)

	log.Debug("processing make execution handler")

	responseHandler.JSONReponse("", http.StatusOK)
}
