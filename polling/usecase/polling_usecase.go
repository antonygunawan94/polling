package usecase

import (
	"log"

	"github.com/antony/polling/polling/model"
	"github.com/antony/polling/polling/repository"
	_pdaRepository "github.com/antony/polling/polling_defined_answer/repository"
)

//

type PollingUsecase interface {
	GetAll() ([]model.Polling, error)
	GetByRoomID(id int64) (*model.Polling, error)
	Insert(polling *model.Polling) error
}

type pollingUsecase struct {
	pollingRepo              repository.PollingRepository
	pollingDefinedAnswerRepo _pdaRepository.PollingDefinedAnswerRepository
}

func (pu *pollingUsecase) GetAll() ([]model.Polling, error) {
	pollings := make([]model.Polling, 0)

	pollings, err := pu.pollingRepo.GetAll()
	if err != nil {
		log.Println(err)
		return pollings, err
	}

	//Create slice references to mutate pollings slice
	pollingsRef := pollings[:0]
	for _, polling := range pollings {
		pdas, err := pu.pollingDefinedAnswerRepo.GetByPollingID(polling.ID)
		if err != nil {
			log.Println(err)
			return pollings, err
		}

		polling.PollingDefinedAnswers = pdas
		pollingsRef = append(pollingsRef, polling)
	}

	return pollings, nil
}

func (pu *pollingUsecase) GetByRoomID(id int64) (*model.Polling, error) {
	polling := &model.Polling{}

	polling, err := pu.pollingRepo.GetByRoomID(id)
	if err != nil {
		log.Println(err)
		return polling, err
	}

	pdas, err := pu.pollingDefinedAnswerRepo.GetByPollingID(polling.ID)
	if err != nil {
		log.Println(err)
		return polling, err
	}

	polling.PollingDefinedAnswers = pdas

	return polling, nil
}

func (pu *pollingUsecase) Insert(polling *model.Polling) error {

	return nil
}

func NewPollingUsecase(pollRepo repository.PollingRepository, pdaRepo _pdaRepository.PollingDefinedAnswerRepository) PollingUsecase {
	return &pollingUsecase{
		pollRepo,
		pdaRepo,
	}
}
