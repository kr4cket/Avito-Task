package models

type Operation struct {
	Id          int    `db:"id"`
	UserId      string `db:"user_id"`
	SegmentName string `db:"segment_name"`
	Action      string `db:"is_delete"`
	Time        string `db:"time"`
}
