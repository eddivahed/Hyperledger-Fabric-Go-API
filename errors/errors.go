// errors/errors.go
package errors

import "net/http"

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewAPIError(status int, message string) *APIError {
	return &APIError{
		Status:  status,
		Message: message,
	}
}

func (e *APIError) Error() string {
	return e.Message
}

func HandleError(w http.ResponseWriter, err error) {
	apiError, ok := err.(*APIError)
	if !ok {
		apiError = NewAPIError(http.StatusInternalServerError, "Internal Server Error")
	}

	w.WriteHeader(apiError.Status)
	w.Write([]byte(apiError.Error()))
}