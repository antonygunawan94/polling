package usecase

import (
	"fmt"
	"log"

	"github.com/antony/polling/polling/model"
	"github.com/antony/polling/polling/repository"
	_pdaRepository "github.com/antony/polling/polling_defined_answer/repository"
	pua "github.com/antony/polling/polling_user_answer/model"
	_puaRepostiory "github.com/antony/polling/polling_user_answer/repository"
)

type PollingUsecase struct {
	pollingRepo              repository.PollingRepository
	pollingDefinedAnswerRepo _pdaRepository.PollingDefinedAnswerRepository
	pollingUserAnswerRepo    _puaRepostiory.PollingUserAnswerRepository
}

func (pu *PollingUsecase) GetAll() ([]model.Polling, error) {
	pollings := make([]model.Polling, 0)

	pollings, err := pu.pollingRepo.GetAll()
	if err != nil {
		log.Println(err)
		return pollings, err
	}

	//Create slice references to mutate pollings slice
	if pollings != nil {
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
	}

	return pollings, nil
}

func (pu *PollingUsecase) GetByID(id int64) (*model.Polling, error) {
	polling := &model.Polling{}

	polling, err := pu.pollingRepo.GetByID(id)
	if err != nil {
		log.Println(err)
		return polling, err
	}

	if polling != nil {
		pdas, err := pu.pollingDefinedAnswerRepo.GetByPollingID(polling.ID)
		if err != nil {
			log.Println(err)
			return polling, err
		}
		polling.PollingDefinedAnswers = pdas
	}

	return polling, nil
}

func (pu *PollingUsecase) GetPollingDetailByID(id int64) (*model.Polling, []pua.PollingUserAnswer, error) {
	polling := &model.Polling{}
	puas := make([]pua.PollingUserAnswer, 0)

	polling, err := pu.pollingRepo.GetByID(id)
	if err != nil {
		log.Println(err)
		return nil, puas, err
	}

	//If polling exist
	if polling != nil {

		//Get all predefined answer for polling
		pdas, err := pu.pollingDefinedAnswerRepo.GetByPollingID(polling.ID)
		if err != nil {
			log.Println(err)
			return nil, puas, err
		}
		polling.PollingDefinedAnswers = pdas

		//Get all user answer for this polling
		puas, err = pu.pollingUserAnswerRepo.GetByPollingID(polling.ID)
		if err != nil {
			log.Println(err)
			return nil, puas, err
		}
	}

	return polling, puas, nil
}

func (pu *PollingUsecase) GetByRoomID(id int64) (*model.Polling, error) {
	polling := &model.Polling{}

	polling, err := pu.pollingRepo.GetByRoomID(id)
	if err != nil {
		log.Println(err)
		return polling, err
	}

	if polling != nil {
		pdas, err := pu.pollingDefinedAnswerRepo.GetByPollingID(polling.ID)
		if err != nil {
			log.Println(err)
			return polling, err
		}

		polling.PollingDefinedAnswers = pdas
	}

	return polling, nil
}

func (pu *PollingUsecase) Insert(polling *model.Polling) error {
	lastID, err := pu.pollingRepo.Insert(polling)
	if err != nil {
		log.Panicln(err)
		return err
	}

	for _, pda := range polling.PollingDefinedAnswers {
		pda.PollingID = lastID
		_, err = pu.pollingDefinedAnswerRepo.Insert(&pda)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (pu *PollingUsecase) AnswerPolling(answer *pua.PollingUserAnswer) error {
	isExist, err := pu.pollingUserAnswerRepo.IsExist(answer)
	if err != nil {
		log.Println(err)
		return err
	}

	if isExist {
		return fmt.Errorf("Username %v has already answer this polling", answer.Username)
	}

	_, err = pu.pollingUserAnswerRepo.Insert(answer)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func NewPollingUsecase(
	pollRepo repository.PollingRepository,
	pdaRepo _pdaRepository.PollingDefinedAnswerRepository,
	puaRepo _puaRepostiory.PollingUserAnswerRepository) *PollingUsecase {

	return &PollingUsecase{
		pollRepo,
		pdaRepo,
		puaRepo,
	}
}
