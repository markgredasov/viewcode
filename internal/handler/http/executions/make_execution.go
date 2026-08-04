package handler_executions

import (
	"net/http"

	"github.com/markgredasov/viewcode/internal/logger"
	"github.com/markgredasov/viewcode/internal/server/http/request"
	"github.com/markgredasov/viewcode/internal/server/http/response"
)

type makeExecutionDTORequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func (h *ExecutionsHTTPHandler) MakeExecution(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, w)

	var req makeExecutionDTORequest
	if err := request.DecodeAndValidate(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "decode and validate")
		return
	}

	result, err := h.executionService.MakeExecution(ctx, req.Language, req.Code)
	if err != nil {
		responseHandler.ErrorResponse(err, "make execution")
		return
	}

	responseHandler.JSONResponse(result, http.StatusOK)
}
