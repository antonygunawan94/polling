package repository

import "github.com/antony/polling/polling_defined_answer/model"

type PollingDefinedAnswerRepository interface {
	GetByPollingID(id int64) ([]model.PollingDefinedAnswer, error)
}
