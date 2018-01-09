package model

import (
	"time"
)

type PollingUserAnswer struct {
	ID                     int64     `json:"id"`
	PollingID              int64     `json:"polling_id"`
	PollingDefinedAnswerID int64     `json:"polling_defined_answer_id"`
	Username               string    `json:"username"`
	Answer                 string    `json:"answer"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}
