package utils

import (
	"github.com/google/uuid"
)

func NewContainerID() string {
	return uuid.NewString()
}
