package polling_defined_answer

import "time"

type PollingDefinedAnswer struct {
	ID        int64     `json:"id"`
	PollingID int64     `json:"polling_id"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
