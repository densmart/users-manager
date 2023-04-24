package services

import (
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/densmart/users-manager/internal/domain/repo"
)

type UsersService struct {
	repo repo.Users
}

func NewUsersService(repo repo.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) Create(data dto.CreateUserDTO) (entities.User, error) {
	return s.repo.Create(data)
}

func (s *UsersService) Update(data dto.UpdateUserDTO) (entities.User, error) {
	return s.repo.Update(data)
}

func (s *UsersService) Retrieve(id uint64) (entities.User, error) {
	return s.repo.Retrieve(id)
}

func (s *UsersService) Search(data dto.SearchUserDTO) ([]entities.User, error) {
	return s.repo.Search(data)
}

func (s *UsersService) Delete(id uint64) error {
	return s.repo.Delete(id)
}
