package error_handler

import (
	"fmt"
	"net/http"
)

// CustomError Error during Authentication
type CustomError struct {
	Msg string
	RedirectUrl string
	Action string
	Code int
}

// Error interface
func (ce *CustomError) Error() string {
	errorString := fmt.Sprintf("msg: [%s] redirectURL: [%s] action: [%s] httpCode: [%d]", ce.Msg, ce.RedirectUrl, ce.Action, ce.Code)
	return errorString
}

// NewCustomError CustomError constructor
func NewCustomError(msg, url, action string, code int) *CustomError {
	return &CustomError{
		Msg: msg,
		RedirectUrl: url,
		Action: action,
		Code: code,
	}
}

// HandleAuthError handles Authentication Error
func HandleAuthError(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/auth-error", 301)
}
