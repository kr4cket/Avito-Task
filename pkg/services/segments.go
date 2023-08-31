package services

import (
	"avitoTask/pkg/models"
	"avitoTask/pkg/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) Create(segment models.Segment) (bool, error) {
	return s.repo.Create(segment)
}

func (s *SegmentService) Delete(segmentName string) (bool, error) {
	return s.repo.Delete(segmentName)
}

func (s *SegmentService) GetAllSegments() ([]models.Segment, error) {
	return s.repo.GetAllSegments()
}

func (s *SegmentService) GetSegmentById(id int) (models.Segment, error) {
	return s.repo.GetSegmentById(id)
}

func (s *SegmentService) CheckSegment(name string) (bool, error) {
	return s.repo.CheckSegment(name)
}
