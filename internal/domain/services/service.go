package services

import "github.com/densmart/users-manager/internal/domain/repo"

type Migrator interface {
	Up() error
	Down() error
}

type Service struct {
	Migrator
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		Migrator: NewMigratorService(repo.Migrator),
	}
}
