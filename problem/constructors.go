package problem

import (
	"net/http"
)

const (
	// DefaultErrorDocumentationLocation is the default URL pointing to the documentation
	// for Problem Details format. It can be overridden using the ErrorDocumentationLocation var.
	DefaultErrorDocumentationLocation = "https://pkg.go.dev/github.com/nickbryan/httputil/problem#"
)

// ErrorDocumentationLocation specifies the URL for the documentation of the Problem Details format.
// This variable can be customized to point to your own API documentation or a different reference.
var ErrorDocumentationLocation = DefaultErrorDocumentationLocation //nolint:gochecknoglobals // Global var improves API without degrading user experience.

// Field represents a specific field that caused a violation constraint.
// It includes details about the error and a pointer to the field in the request body.
type Field struct {
	Detail  string `json:"detail"`
	Pointer string `json:"pointer"`
}

// BadRequest creates a DetailedError for bad request errors.
func BadRequest(r *http.Request) *DetailedError {
	return &DetailedError{
		Type:             DefaultErrorDocumentationLocation + "BadRequest",
		Title:            "Bad Request",
		Detail:           "The request is invalid or malformed",
		Status:           http.StatusBadRequest,
		Instance:         r.URL.Path,
		ExtensionMembers: nil,
	}
}

// ConstraintViolation creates a DetailedError for constraint violation errors.
// The Field describe the specific fields that violated constraints.
func ConstraintViolation(r *http.Request, fields ...Field) *DetailedError {
	return &DetailedError{
		Type:             DefaultErrorDocumentationLocation + "ConstraintViolation",
		Title:            "Constraint Violation",
		Detail:           "The request data violated one or more validation constraints",
		Status:           http.StatusBadRequest,
		Instance:         r.URL.Path,
		ExtensionMembers: map[string]any{"violations": fields},
	}
}

// ServerError creates a DetailedError for internal server errors.
func ServerError(r *http.Request) *DetailedError {
	return &DetailedError{
		Type:             DefaultErrorDocumentationLocation + "ServerError",
		Title:            "Server Error",
		Detail:           "The server encountered an unexpected internal error",
		Status:           http.StatusInternalServerError,
		Instance:         r.URL.Path,
		ExtensionMembers: nil,
	}
}