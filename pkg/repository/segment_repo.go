package repository

import (
	"avitoTask/pkg/models"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (r *SegmentPostgres) Create(segment models.Segment) (bool, error) {

	_, err := r.db.Exec("INSERT INTO segments (name) VALUES ($1)", segment.Name)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *SegmentPostgres) Delete(name string) (bool, error) {

	_, err := r.db.Exec("DELETE from segments WHERE name = $1", name)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *SegmentPostgres) GetAllSegments() ([]models.Segment, error) {

	var segments []models.Segment

	err := r.db.Select(&segments, "SELECT * FROM segments")

	return segments, err
}

func (r *SegmentPostgres) GetSegmentById(id int) (models.Segment, error) {

	var segment models.Segment

	err := r.db.Get(&segment, "SELECT * FROM segments WHERE id = $1", id)

	return segment, err
}

func (r *SegmentPostgres) CheckSegment(segmentName string) (bool, error) {

	var checkTable string

	fmt.Println(segmentName)

	row := r.db.QueryRow("SELECT name FROM segments WHERE name = $1", segmentName)

	err := row.Scan(&checkTable)

	fmt.Println(err)

	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
