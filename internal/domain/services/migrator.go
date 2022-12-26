package services

import "github.com/densmart/users-manager/internal/domain/repo"

type MigratorService struct {
	repo repo.Migrator
}

func NewMigratorService(repo repo.Migrator) *MigratorService {
	return &MigratorService{
		repo: repo,
	}
}

func (s *MigratorService) Up() error {
	return s.repo.Up()
}

func (s *MigratorService) Down() error {
	return s.repo.Down()
}
