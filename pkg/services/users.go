package services

import (
	"avitoTask/pkg/models"
	"avitoTask/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddUser(user models.User) (bool, error) {
	return s.repo.AddUser(user)
}

func (s *UserService) ChangeSegments(userId int, segmentsToAdd []string, segmentsToDelete []string) (bool, error) {
	return s.repo.ChangeSegments(userId, segmentsToAdd, segmentsToDelete)
}

func (s *UserService) GetActiveUserSegments(userId int) ([]string, error) {
	return s.repo.GetActiveUserSegments(userId)
}

func (s *UserService) SetExpiredSegment(userId int, ttl int, segment string) (bool, error) {
	return s.repo.SetExpiredSegment(userId, ttl, segment)
}

func (s *UserService) DeleteExpiredSegments(userId int) ([]string, error) {
	return s.repo.DeleteExpiredSegments(userId)
}

func (s *UserService) GetRandomUsers(entirety int16) []int {
	return s.repo.GetRandomUsers(entirety)
}
