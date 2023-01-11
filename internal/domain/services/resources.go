package services

import (
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/densmart/users-manager/internal/domain/repo"
)

type ResourcesService struct {
	repo repo.Resources
}

func NewResourcesService(repo repo.Resources) *ResourcesService {
	return &ResourcesService{repo: repo}
}

func (s *ResourcesService) Create(data dto.CreateResourceDTO) (entities.Resource, error) {
	return s.repo.Create(data)
}

func (s *ResourcesService) Update(data dto.UpdateResourceDTO) (entities.Resource, error) {
	return s.repo.Update(data)
}

func (s *ResourcesService) Retrieve(id uint64) (entities.Resource, error) {
	return s.repo.Retrieve(id)
}

func (s *ResourcesService) Search(data dto.SearchResourceDTO) ([]entities.Resource, error) {
	return s.repo.Search(data)
}

func (s *ResourcesService) Delete(id uint64) error {
	return s.repo.Delete(id)
}
