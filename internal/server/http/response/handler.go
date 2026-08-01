package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/markgredasov/viewcode/internal/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *logger.Logger
	w   http.ResponseWriter
}

func NewHTTPResponseHandler(l *logger.Logger, w http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: l,
		w:   w,
	}
}

func (h *HTTPResponseHandler) JSONReponse(responseBody any, statusCode int) {
	h.w.WriteHeader(statusCode)

	if err := json.NewEncoder(h.w).Encode(responseBody); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))

	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))

	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) errorResponse(statusCode int, err error, msg string) {
	errorResponse := ErrorResponse{
		Error:   err,
		Message: msg,
	}

	h.JSONReponse(errorResponse, statusCode)
}
