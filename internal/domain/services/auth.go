package services

import "github.com/spf13/viper"

type AuthService struct {
	currentAPIKey     string
	disconnectedUsers map[uint64]bool
}

func NewAuthService() *AuthService {
	return &AuthService{
		disconnectedUsers: make(map[uint64]bool),
	}
}

func (s *AuthService) GetAPIKey() string {
	if s.currentAPIKey == "" {
		// @TODO: get key from some storage
		s.currentAPIKey = viper.GetString("app.auth_key")
	}
	return s.currentAPIKey
}

func (s *AuthService) RenewAPIKey() {
	s.currentAPIKey = ""
	s.GetAPIKey()
}

func (s *AuthService) CleanAPIKey() {
	if s.currentAPIKey != "" {
		s.currentAPIKey = ""
	}
}

func (s *AuthService) AddDisconnectedUser(id uint64) {
	_, ok := s.disconnectedUsers[id]
	if !ok {
		s.disconnectedUsers[id] = true
	}
}

func (s *AuthService) CheckDisconnectedUser(id uint64) bool {
	_, ok := s.disconnectedUsers[id]
	return ok
}

func (s *AuthService) RemoveDisconnectedUser(id uint64) {
	_, ok := s.disconnectedUsers[id]
	if ok {
		delete(s.disconnectedUsers, id)
	}
}
