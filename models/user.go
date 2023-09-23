package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id           uuid.UUID
	Name         *string
	Login        *string
	Password     *string
	Created      time.Time
	Token        *string
	RefreshToken *string
}
