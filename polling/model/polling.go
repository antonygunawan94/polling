package model

import (
	"time"

	pda "github.com/antony/polling/polling_defined_answer/model"
)

type Polling struct {
	ID                    int64                      `json:"id"`
	RoomID                int64                      `json:"room_id"`
	Question              string                     `json:"question"`
	PollingDefinedAnswers []pda.PollingDefinedAnswer `json:"polling_defined_answers"`
	StartDate             time.Time                  `json:"start_date"`
	ExpiredDate           time.Time                  `json:"expired_date"`
	CreatedAt             time.Time                  `json:"created_at"`
	UpdatedAt             time.Time                  `json:"updated_at"`
}
