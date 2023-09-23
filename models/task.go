package models

import (
	"github.com/google/uuid"
	"time"
)

type EnumStatus int

const (
	New EnumStatus = iota
	InProgress
	Done
	Overdue
)

type Task struct {
	Id          uuid.UUID
	Title       *string
	Description *string
	Status      EnumStatus
	Created     time.Time
	Checked     time.Time
	User_id     uuid.UUID
}
