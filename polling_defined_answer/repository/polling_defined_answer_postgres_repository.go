package repository

import (
	"log"

	"github.com/antony/polling/polling_defined_answer/model"
	"github.com/tokopedia/sqlt"
)

type postgresPollingDefinedAnswerRepository struct {
	Conn *sqlt.DB
}

func NewPostgresPollingDefinedAnswerRepository(conn *sqlt.DB) PollingDefinedAnswerRepository {
	return &postgresPollingDefinedAnswerRepository{conn}
}

func (ppdar *postgresPollingDefinedAnswerRepository) GetByPollingID(id int64) ([]model.PollingDefinedAnswer, error) {
	query := `
			SELECT 
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
		return pdas, err
	}
	for rows.Next() {
		pda := model.PollingDefinedAnswer{}
		err = rows.Scan(&pda.ID, &pda.PollingID, &pda.Answer,
			&pda.CreatedAt, &pda.UpdatedAt)
		pdas = append(pdas, pda)
	}
	return pdas, nil
}

func (ppdar *postgresPollingDefinedAnswerRepository) Insert(pda *model.PollingDefinedAnswer) (int64, error) {
	query := `
			INSERT INTO
				polling_defined_answers
				(
					polling_id,
					answer,
					created_at,
					updated_at	
				)
			VALUES
				(
					$1,
					$2,
					now(),
					now()
				)
			RETURNING id
	`
	var lastID int64
	err := ppdar.Conn.QueryRow(query, pda.PollingID, pda.Answer).Scan(&lastID)
	if err != nil {
		log.Println(err)
		return lastID, err
	}

	return lastID, nil
}

func (ppdar *postgresPollingDefinedAnswerRepository) Update(pda *model.PollingDefinedAnswer) error {
	query := `
		UPDATE
			polling_defined_answers
		SET
			polling_id = $1,
			answer = $2,
			updated_at = now()
		WHERE
			id = $3
	`

	stmt, err := ppdar.Conn.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pda.PollingID, pda.Answer, pda.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (ppdar *postgresPollingDefinedAnswerRepository) Delete(id int64) error {
	query := `
		DELETE
			polling_defined_answers
		WHERE
			id = $1
	`

	stmt, err := ppdar.Conn.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
