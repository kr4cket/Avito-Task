package services

import (
	"avitoTask/pkg/models"
	"avitoTask/pkg/repository"
)

type User interface {
	AddUser(user models.User) (bool, error)
	ChangeSegments(userId int, segmentsToAdd []string, segmentsToDelete []string) (bool, error)
	GetActiveUserSegments(userId int) ([]string, error)
	SetExpiredSegment(userId int, ttl int, segment string) (bool, error)
	DeleteExpiredSegments(userId int) ([]string, error)
	GetRandomUsers(entirety int16) []int
}

type Segment interface {
	Create(segments models.Segment) (bool, error)
	Delete(name string) (bool, error)
	GetSegmentById(id int) (models.Segment, error)
	GetAllSegments() ([]models.Segment, error)
	CheckSegment(name string) (bool, error)
}

type Operation interface {
	AddHistoryRecord(userId int, segmentsToAdd []string, segmentsToDelete []string)
	GetHistory(userId int, month int, year int) ([]models.Operation, error)
}

type Service struct {
	User
	Segment
	Operation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Segment:   NewSegmentService(repos.Segment),
		User:      NewUserService(repos.User),
		Operation: NewOperationService(repos.Operation),
	}
}
