package services

import (
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/densmart/users-manager/internal/domain/repo"
)

type ActionsService struct {
	repo repo.Actions
}

func NewActionsService(repo repo.Actions) *ActionsService {
	return &ActionsService{repo: repo}
}

func (s *ActionsService) Create(data dto.CreateActionDTO) (entities.Action, error) {
	return s.repo.Create(data)
}

func (s *ActionsService) Update(data dto.UpdateActionDTO) (entities.Action, error) {
	return s.repo.Update(data)
}

func (s *ActionsService) Retrieve(id uint64) (entities.Action, error) {
	return s.repo.Retrieve(id)
}

func (s *ActionsService) Search(data dto.SearchActionDTO) ([]entities.Action, error) {
	return s.repo.Search(data)
}

func (s *ActionsService) Delete(id uint64) error {
	return s.repo.Delete(id)
}
