package repository

import (
	"log"

	model "github.com/antony/polling/polling_defined_answer"
	"github.com/tokopedia/sqlt"
)

type postgresPollingDefinedAnswerRepository struct {
	Conn *sqlt.DB
}

func NewPostgresPollingDefinedAnswerRepository(conn *sqlt.DB) PollingDefinedAnswerRepository {
	return &postgresPollingDefinedAnswerRepository{conn}
}

func (ppdar *postgresPollingDefinedAnswerRepository) GetByPollingID(id int64) (*[]model.PollingDefinedAnswer, error) {
	query := `SELECT 
				id,
				polling_id,
				answer,
				created_at,
				updated_at
			FROM
				polling_defined_answers
			WHERE
				polling_id = $1`

	pdas := make([]model.PollingDefinedAnswer, 0)
	rows, err := ppdar.Conn.Query(query, id)
	if err != nil {
		log.Println(err)
		return &pdas, err
	}
	for rows.Next() {
		pda := model.PollingDefinedAnswer{}
		err = rows.Scan(&pda.ID, &pda.PollingID, &pda.Answer,
			&pda.CreatedAt, &pda.UpdatedAt)
		pdas = append(pdas, pda)
	}
	return &pdas, nil
}
