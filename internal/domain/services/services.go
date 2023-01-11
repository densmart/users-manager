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

type Actions interface {
	Create(data dto.CreateActionDTO) (entities.Action, error)
	Update(data dto.UpdateActionDTO) (entities.Action, error)
	Retrieve(id uint64) (entities.Action, error)
	Search(data dto.SearchActionDTO) ([]entities.Action, error)
	Delete(id uint64) error
}

type Resources interface {
	Create(data dto.CreateResourceDTO) (entities.Resource, error)
	Update(data dto.UpdateResourceDTO) (entities.Resource, error)
	Retrieve(id uint64) (entities.Resource, error)
	Search(data dto.SearchResourceDTO) ([]entities.Resource, error)
	Delete(id uint64) error
}

type Service struct {
	Migrator
	Roles
	Actions
	Resources
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		Migrator:  NewMigratorService(repo.Migrator),
		Roles:     NewRolesService(repo.Roles),
		Actions:   NewActionsService(repo.Actions),
		Resources: NewResourcesService(repo.Resources),
	}
}
