package services

import (
	"avitoTask/pkg/models"
	"avitoTask/pkg/repository"
)

type OperationService struct {
	repo repository.Operation
}

func NewOperationService(repo repository.Operation) *OperationService {
	return &OperationService{repo: repo}
}

func (s *OperationService) AddHistoryRecord(userId int, segmentsToAdd []string, segmentsToDelete []string) {
	s.repo.AddHistoryRecord(userId, segmentsToAdd, segmentsToDelete)
}

func (s *OperationService) GetHistory(userId int, month int, year int) ([]models.Operation, error) {
	return s.repo.GetHistory(userId, month, year)
}
