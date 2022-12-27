package services

import (
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/densmart/users-manager/internal/domain/repo"
)

type RolesService struct {
	repo repo.Roles
}

func NewRolesService(repo repo.Roles) *RolesService {
	return &RolesService{repo: repo}
}

func (s *RolesService) Create(data dto.CreateRoleDTO) (entities.Role, error) {
	return s.repo.Create(data)
}

func (s *RolesService) Update(data dto.UpdateRoleDTO) (entities.Role, error) {
	return s.repo.Update(data)
}

func (s *RolesService) Retrieve(id uint64) (entities.Role, error) {
	return s.repo.Retrieve(id)
}

func (s *RolesService) Search(data dto.SearchRoleDTO) ([]entities.Role, error) {
	return s.repo.Search(data)
}

func (s *RolesService) Delete(id uint64) error {
	return s.repo.Delete(id)
}
