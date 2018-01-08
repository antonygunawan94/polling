package repository

import (
	"database/sql"
	"log"

	model "github.com/antony/polling/polling"
	"github.com/tokopedia/sqlt"
)

type postgresPollingRepository struct {
	Conn *sqlt.DB
}

func NewPostgresPollingRepository(Conn *sqlt.DB) PollingRepository {
	return &postgresPollingRepository{Conn}
}

func (ppr *postgresPollingRepository) GetAll() ([]model.Polling, error) {
	query := `
			SELECT 
				id, 
				room_id, 
				question, 
				start_date, 
				expired_date, 
				created_at,
				updated_at
			FROM
				pollings
			`
	pollings := make([]model.Polling, 0)

	rows, err := ppr.Conn.Query(query)
	if err != nil {
		log.Println(err)
		return pollings, err
	}
	defer rows.Close()

	for rows.Next() {
		polling := model.Polling{}
		err = rows.Scan(
			&polling.ID,
			&polling.RoomID,
			&polling.Question,
			&polling.StartDate,
			&polling.ExpiredDate,
			&polling.CreatedAt,
			&polling.UpdatedAt,
		)
		if err != nil{
			log.Println(err)
			return pollings, err
		}

		pollings = append(pollings, polling)
	}
	
	return pollings, nil
}

func (ppr *postgresPollingRepository) GetByRoomID(id int64) (*model.Polling, error) {
	query := `
			SELECT 
				id, 
				room_id, 
				question, 
				start_date, 
				expired_date, 
				created_at,
				updated_at
			FROM
				pollings
			WHERE
				room_id = $1
			LIMIT 1
			`

	polling := model.Polling{}

	err := ppr.Conn.QueryRow(query, id).Scan(&polling.ID, &polling.RoomID,
		&polling.Question, &polling.StartDate, &polling.ExpiredDate,
		&polling.CreatedAt, &polling.UpdatedAt)

	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return &polling, err
	}

	return &polling, nil
}

func (ppr *postgresPollingRepository) Insert(polling *model.Polling) error {
	query := `
			INSERT INTO pollings (
					room_id,
					question,
					start_date,
					expired_date,
					created_at,
					updated_at
			) VALUES (
					$1,
					$2,
					$3,
					$4,
					now(),
					now()
			)
			`
	stmt, err := ppr.Conn.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(polling.RoomID, polling.Question, polling.StartDate, polling.ExpiredDate)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ppr *postgresPollingRepository) Update(polling *model.Polling) error {
	query := `
		UPDATE 
			pollings
		SET
			room_id = $1,
			question = $2,
			start_date = $3,
			expired_date = $4,
			update_at = now()
		WHERE
			id = $5
	`

	stmt, err := ppr.Conn.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(polling.RoomID, polling.Question, polling.StartDate, polling.ExpiredDate, polling.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (ppr *postgresPollingRepository) Delete(id int64) error {
	query := `
		DELETE
			pollings
		WHERE
			id = $1
	`
	stmt, err := ppr.Conn.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
