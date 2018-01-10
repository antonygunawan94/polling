package repository

import "github.com/antony/polling/polling_user_answer/model"

type PollingUserAnswerRepository interface {
	GetByPollingID(id int64) ([]model.PollingUserAnswer, error)
	Insert(answer *model.PollingUserAnswer) (int64, error)
	IsExist(answer *model.PollingUserAnswer) (bool, error)
}
