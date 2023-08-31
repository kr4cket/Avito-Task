package repository

import (
	"avitoTask/pkg/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type OperationPostgres struct {
	db *sqlx.DB
}

func NewOperationPostgres(db *sqlx.DB) *OperationPostgres {
	return &OperationPostgres{db: db}
}

func (r *OperationPostgres) AddHistoryRecord(userId int, segmentsToAdd []string, segmentsToDelete []string) {

	if len(segmentsToDelete) > 0 {
		r.db.Exec("INSERT INTO operation_history (user_id, segment_name, is_delete) VALUES ($1, unnest($2::text[]), true)", userId, pq.Array(segmentsToDelete))
	}

	if len(segmentsToAdd) > 0 {
		r.db.Exec("INSERT INTO operation_history (user_id, segment_name, is_delete) VALUES ($1, unnest($2::text[]), false)", userId, pq.Array(segmentsToAdd))
	}
}

func (r *OperationPostgres) GetHistory(userId int, month int, year int) ([]models.Operation, error) {

	var operations []models.Operation

	err := r.db.Select(&operations, "SELECT * FROM operation_history WHERE user_id = $1 AND extract(year FROM time) = $2 AND extract(month FROM time) = $3 ORDER BY id DESC", userId, year, month)

	return operations, err
}
