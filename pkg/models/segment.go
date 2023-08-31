package models

import "database/sql"

type Segment struct {
	Id       int           `json:"id" db:"id"`
	Name     string        `json:"name" db:"name" binding:"required"`
	Entirety sql.NullInt16 `json:"entirety" db:"entirety"`
}
