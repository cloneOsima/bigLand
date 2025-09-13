package utils

import (
	"github.com/google/uuid"
)

// utils package provides helper functions and common utilities used across the project.
func GenerateRequestId() uuid.UUID {
	reqId, err := uuid.NewRandom()
	if err != nil {
		return reqId
	}

	return reqId
}
