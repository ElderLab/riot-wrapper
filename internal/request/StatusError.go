package request

import "strconv"

// NewStatusError is a function that returns a new StatusError.
func NewStatusError(status int) *StatusError {
	return &StatusError{status: status}
}

// StatusError is a struct that contains the status.
type StatusError struct {
	status int
}

// Error is a function that returns the status.
func (e StatusError) Error() string {
	return "Status : " + strconv.Itoa(e.status)
}
