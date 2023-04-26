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

type Resources interface {
	Create(data dto.CreateResourceDTO) (entities.Resource, error)
	Update(data dto.UpdateResourceDTO) (entities.Resource, error)
	Retrieve(id uint64) (entities.Resource, error)
	Search(data dto.SearchResourceDTO) ([]entities.Resource, error)
	Delete(id uint64) error
}

type Users interface {
	Create(data dto.CreateUserDTO) (entities.User, error)
	Update(data dto.UpdateUserDTO) (entities.User, error)
	Retrieve(id uint64) (entities.User, error)
	Search(data dto.SearchUserDTO) ([]entities.User, error)
	Delete(id uint64) error
}

type Permissions interface {
	Create(data []dto.CreatePermissionDTO) ([]entities.Permission, error)
	Update(data []dto.UpdatePermissionDTO) ([]entities.Permission, error)
	Search(data dto.SearchPermissionDTO) ([]entities.Permission, error)
	Delete(id []uint64) error
}

type Service struct {
	Migrator
	Roles
	Resources
	Users
	Permissions
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		Migrator:    NewMigratorService(repo.Migrator),
		Roles:       NewRolesService(repo.Roles),
		Resources:   NewResourcesService(repo.Resources),
		Users:       NewUsersService(repo.Users),
		Permissions: NewPermissionsService(repo.Permissions),
	}
}
