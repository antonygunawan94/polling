package repository

import (
	"log"

	"github.com/antony/polling/polling_user_answer/model"
	"github.com/tokopedia/sqlt"
)

type pollingUserAnswerPostgresRepository struct {
	conn *sqlt.DB
}

func NewPollingUserAnswerPostgresRepository(conn *sqlt.DB) PollingUserAnswerRepository {
	return &pollingUserAnswerPostgresRepository{
		conn,
	}
}

func (puapr *pollingUserAnswerPostgresRepository) GetByPollingID(id int64) ([]model.PollingUserAnswer, error) {
	query := `
		SELECT 
			pua.id,
			pua.polling_id,
			pua.polling_defined_answer_id,
			pua.username,
			COALESCE(pda.answer, pua.custom_answer) as answer,
			pua.created_at,
			pua.updated_at
		FROM 
			polling_user_answers pua
		LEFT JOIN
			polling_defined_answers pda
		ON
			pua.polling_defined_answer_id = pda.id
		WHERE
			pua.polling_id = $1
	`

	puas := make([]model.PollingUserAnswer, 0)

	rows, err := puapr.conn.Query(query, id)
	if err != nil {
		log.Println(err)
		return puas, err
	}
	defer rows.Close()

	for rows.Next() {
		pua := model.PollingUserAnswer{}
		err = rows.Scan(
			&pua.ID,
			&pua.PollingID,
			&pua.PollingDefinedAnswerID,
			&pua.Username,
			&pua.Answer,
			&pua.CreatedAt,
			&pua.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return puas, err
		}
		puas = append(puas, pua)
	}

	return puas, nil
}

func (puapr *pollingUserAnswerPostgresRepository) Insert(answer *model.PollingUserAnswer) (int64, error) {
	query := `
		INSERT INTO
			polling_user_answers
		(
			polling_id,
			poling_defined_answer_id,
			username,
			custom_answer,
			created_at,
			updated_at
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4,
			now(),
			now()
		)
		RETURNING id
	`
	var lastID int64 = -1
	err := puapr.conn.QueryRow(
		query,
		answer.PollingID,
		answer.PollingDefinedAnswerID,
		answer.Username,
		answer.Answer,
	).Scan(&lastID)

	if err != nil {
		log.Println(err)
		return lastID, err
	}
	return lastID, nil
}
