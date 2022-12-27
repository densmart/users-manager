package services

import (
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/densmart/users-manager/internal/domain/repo"
)

type Migrator interface {
	Up() error
	Down() error
}

type Roles interface {
	Create(data dto.CreateRoleDTO) (entities.Role, error)
	Update(data dto.UpdateRoleDTO) (entities.Role, error)
	Retrieve(id uint64) (entities.Role, error)
	Search(data dto.SearchRoleDTO) ([]entities.Role, error)
	Delete(id uint64) error
}

type Service struct {
	Migrator
	Roles
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		Migrator: NewMigratorService(repo.Migrator),
		Roles:    NewRolesService(repo.Roles),
	}
}
