package errors

// NewRateLimitedError is a function that returns a new RateLimitedError.
func NewRateLimitedError() *RateLimitedError {
	return &RateLimitedError{}
}

// RateLimitedError is a struct that contains the status.
type RateLimitedError struct{}

// Error is a function that returns the status.
func (e RateLimitedError) Error() string {
	return "Rate Limited"
}
