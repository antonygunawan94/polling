package repository

import "github.com/antony/polling/polling/model"

type PollingRepository interface {
	GetAll() ([]model.Polling, error)
	GetByID(id int64) (*model.Polling, error)
	GetByRoomID(id int64) (*model.Polling, error)
	Insert(polling *model.Polling) (int64, error)
	Update(polling *model.Polling) error
	Delete(id int64) error
}
