package services

import (
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/densmart/users-manager/internal/domain/repo"
)

type PermissionsService struct {
	repo repo.Permissions
}

func NewPermissionsService(repo repo.Permissions) *PermissionsService {
	return &PermissionsService{repo: repo}
}

func (s *PermissionsService) Create(data []dto.CreatePermissionDTO) ([]entities.Permission, error) {
	return s.repo.Create(data)
}

func (s *PermissionsService) Update(data []dto.UpdatePermissionDTO) ([]entities.Permission, error) {
	return s.repo.Update(data)
}

func (s *PermissionsService) Search(data dto.SearchPermissionDTO) ([]entities.Permission, error) {
	return s.repo.Search(data)
}

func (s *PermissionsService) Delete(id []uint64) error {
	return s.repo.Delete(id)
}
