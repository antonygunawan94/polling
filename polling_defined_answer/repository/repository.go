package repository

import (
	model "github.com/antony/polling/polling_defined_answer"
)

type PollingDefinedAnswerRepository interface {
	GetByPollingID(id int64) (*[]model.PollingDefinedAnswer, error)
}
