package httperr

import (
	"net/http"
)

type HTTPError struct {
	err  error
	code int
}

func New(err error, statusCode int) *HTTPError {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}

	return &HTTPError{err, statusCode}
}

func (h *HTTPError) StatusCode() int {
	return h.code
}

func (h *HTTPError) Error() string {
	return h.err.Error()
}

func (h *HTTPError) Write(res http.ResponseWriter) {
	http.Error(res, h.err.Error(), h.code)
}
