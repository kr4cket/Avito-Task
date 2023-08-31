package repository

import (
	"avitoTask/pkg/models"
	"crypto/sha1"
	"fmt"
	"math"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash)
}

func (r *UserPostgres) AddUser(user models.User) (bool, error) {

	pass := generatePasswordHash(user.Password)

	_, err := r.db.Exec("INSERT INTO users (login, password) VALUES ($1, $2)", user.Login, pass)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserPostgres) ChangeSegments(userId int, segmentsToAdd []string, segmentsToDelete []string) (bool, error) {

	if len(segmentsToDelete) > 0 {

		_, err := r.db.Exec("DELETE FROM user_segments WHERE user_id = $1 and segment_name IN (SELECT unnest($2::text[]))", userId, pq.Array(segmentsToDelete))

		if err != nil {
			return false, err
		}

	}

	if len(segmentsToAdd) > 0 {

		_, err := r.db.Exec("INSERT INTO user_segments (user_id, segment_name) VALUES ($1, unnest($2::text[]))", userId, pq.Array(segmentsToAdd))

		if err != nil {
			return false, err
		}

	}

	return true, nil
}

func (r *UserPostgres) GetActiveUserSegments(userId int) ([]string, error) {

	var segments []string

	err := r.db.Select(&segments, "SELECT segment_name FROM user_segments WHERE user_id = $1", userId)

	return segments, err
}

func (r *UserPostgres) SetExpiredSegment(userId int, ttl int, segment string) (bool, error) {

	expires := fmt.Sprintf("%d hours", ttl)

	_, err := r.db.Exec("INSERT INTO user_segments (user_id, segment_name, expire) VALUES ($1, $2, (current_timestamp + $3::interval))", userId, segment, expires)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserPostgres) DeleteExpiredSegments(userId int) ([]string, error) {

	var segments []string

	rows, err := r.db.Query("DELETE FROM user_segments WHERE user_id=$1 and expire IS NOT NULL and expire < CURRENT_TIMESTAMP RETURNING segment_name", userId)

	if err != nil {
		return []string{}, err
	}

	for rows.Next() {
		var segment string
		if err := rows.Scan(&segment); err != nil {
			logrus.Panic("Scan error!")
		}
		segments = append(segments, segment)
	}

	return segments, nil
}

func (r *UserPostgres) GetRandomUsers(entirety int16) []int {
	var users []int
	var count []float64

	r.db.Select(&count, "SELECT COUNT(*) FROM users")

	userCount := math.Ceil(count[0] * (float64(entirety) / 100.0))

	r.db.Select(&users, "SELECT id FROM users ORDER BY RANDOM() LIMIT $1", userCount)

	return users
}
