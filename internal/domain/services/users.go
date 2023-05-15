package services

import (
	"errors"
	"fmt"
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/densmart/users-manager/internal/domain/repo"
	"github.com/densmart/users-manager/internal/logger"
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

func (s *UsersService) RetrieveByEmail(email string) (entities.User, error) {
	searchDTO := dto.SearchUserDTO{
		Email: &email,
	}
	users, err := s.Search(searchDTO)
	if err != nil {
		return entities.User{}, err
	}
	if len(users) > 1 {
		logger.Errorf("find more than one user with email: %s", email)
		return entities.User{}, errors.New(fmt.Sprintf("find more than one user with email: %s", email))
	}
	return users[0], err
}

func (s *UsersService) Search(data dto.SearchUserDTO) ([]entities.User, error) {
	return s.repo.Search(data)
}

func (s *UsersService) Delete(id uint64) error {
	return s.repo.Delete(id)
}
