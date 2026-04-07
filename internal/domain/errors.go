package domain

import "errors"

var (
	ErrTaskNotFound   = errors.New("task not found")
	ErrInvalidStatus  = errors.New("invalid task status")
	ErrInvalidTaskID  = errors.New("invalid task id")
	ErrUnauthorized   = errors.New("unauthorized access to task")
	ErrInvalidPayload = errors.New("invalid request payload")
)
