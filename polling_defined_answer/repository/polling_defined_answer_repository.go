package repository

import "github.com/antony/polling/polling_defined_answer/model"

type PollingDefinedAnswerRepository interface {
	GetByPollingID(id int64) ([]model.PollingDefinedAnswer, error)
	Insert(pda *model.PollingDefinedAnswer) (int64, error)
	Update(pda *model.PollingDefinedAnswer) error
	Delete(id int64) error
}
