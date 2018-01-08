package repository

import (
	model "github.com/antony/polling/polling"
)

type PollingRepository interface {
	GetAll() ([]model.Polling, error)
	GetByRoomID(id int64) (*model.Polling, error)
	Insert(polling *model.Polling) error
	Update(polling *model.Polling) error
	Delete(id int64) error
}
