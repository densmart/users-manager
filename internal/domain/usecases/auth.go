package usecases

import (
	"errors"
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/services"
	"github.com/densmart/users-manager/internal/domain/utils"
	"github.com/densmart/users-manager/internal/domain/utils/jwt_auth"
	"time"
)

func SignIn(s services.Service, data dto.AuthRequestDTO) (*dto.AuthResponseDTO, error) {
	apiKey := s.Auth.GetAPIKey()
	user, err := s.Users.RetrieveByEmail(data.Login)
	if err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, errors.New("inactive user")
	}
	if user.Is2fa {
		return &dto.AuthResponseDTO{
			Need2Fa: true,
		}, nil
	}
	// Check user PWD
	pwdChecked := utils.CheckPasswordHash(user.Password, data.Password)
	if !pwdChecked {
		return nil, errors.New("password is incorrect")
	}
	// Get all role permissions and create permission map like "uri_mask:method_mask"
	permissions, err := getRolePermissions(s.Permissions, user.RoleID)
	if err != nil {
		return nil, err
	}
	// Create JWTs
	jwtWrapper := jwt_auth.NewJwtAuth(apiKey, user.Id, user.FirstName+" "+user.LastName, user.Email, permissions)
	at, rt, err := jwtWrapper.GenerateTokens()
	if err != nil {
		return nil, err
	}
	// update user's last login timestamp
	if err = updateUserLastLogin(s.Users, user.Id); err != nil {
		return nil, err
	}
	// remove user from disconnected list if it's in it
	s.Auth.RemoveDisconnectedUser(user.Id)

	response := dto.AuthResponseDTO{
		Need2Fa:      false,
		AccessToken:  at,
		RefreshToken: rt,
	}

	return &response, nil
}

func RefreshAccessToken(s services.Service, data dto.RefreshRequestDTO) (*dto.RefreshResponseDTO, error) {
	apiKey := s.Auth.GetAPIKey()
	// Check refresh token and receive claims
	jwtToken := jwt_auth.NewJwtToken(apiKey)
	jwtToken.Refresh = data.RefreshToken
	claims, err := jwtToken.GetRefreshClaims()
	if err != nil {
		return nil, err
	}

	// Get user data
	user, err := s.Users.Retrieve(claims.UserID)
	if err != nil {
		return nil, err
	}
	// if user already disconnected - no need to refresh token
	if s.Auth.CheckDisconnectedUser(user.Id) {
		return nil, errors.New("user disconnected, please make re-login")
	}
	// Get all role permissions and create permission map like "uri_mask:method_mask"
	permSearchDTO := dto.SearchPermissionDTO{
		RoleID: user.RoleID,
	}
	perms, err := s.Permissions.Search(permSearchDTO)
	if err != nil {
		return nil, err
	}
	permissions := make(map[string]uint8)
	for _, perm := range perms {
		permissions[perm.ResourceURIMask] = perm.MethodMask
	}
	// Create JWTs
	jwtWrapper := jwt_auth.NewJwtAuth(apiKey, user.Id, user.FirstName+" "+user.LastName, user.Email, permissions)
	accessToken := jwtWrapper.GenerateAccessToken()
	at, err := accessToken.SignedString([]byte(apiKey))
	if err != nil {
		return nil, err
	}
	// Make result
	response := dto.RefreshResponseDTO{
		AccessToken: at,
	}
	return &response, nil
}

func CheckAuthorization(s services.Service, accessToken string) (*dto.AuthDataDTO, error) {
	apiKey := s.Auth.GetAPIKey()
	// Check access token and receive claims
	jwtToken := jwt_auth.NewJwtToken(apiKey)
	jwtToken.Access = accessToken
	claims, err := jwtToken.GetAccessClaims()
	if err != nil {
		return nil, err
	}
	if s.Auth.CheckDisconnectedUser(claims.UserID) {
		return nil, errors.New("user has been disconnected")
	}

	// Make result
	result := dto.AuthDataDTO{
		UserID:      claims.UserID,
		UserName:    claims.UserName,
		Email:       claims.Email,
		Permissions: claims.Permissions,
	}
	return &result, nil
}

func CheckOTPToken(s services.Service, data dto.OTPRequestDTO) (*dto.AuthResponseDTO, error) {
	apiKey := s.Auth.GetAPIKey()
	user, err := s.Users.RetrieveByEmail(data.Login)
	if err != nil {
		return nil, err
	}
	// Verify OTP key
	if err = utils.VerifyOTP(data.OtpKey, user.Token2fa.String); err != nil {
		return nil, err
	}
	// Check user PWD
	pwdChecked := utils.CheckPasswordHash(user.Password, data.Password)
	if !pwdChecked {
		return nil, errors.New("password is incorrect")
	}
	// Get all role permissions and create permission map like "uri_mask:method_mask"
	permissions, err := getRolePermissions(s.Permissions, user.RoleID)
	if err != nil {
		return nil, err
	}
	// Create JWTs
	jwtWrapper := jwt_auth.NewJwtAuth(apiKey, user.Id, user.FirstName+" "+user.LastName, user.Email, permissions)
	at, rt, err := jwtWrapper.GenerateTokens()
	if err != nil {
		return nil, err
	}
	// update user's last login timestamp
	if err = updateUserLastLogin(s.Users, user.Id); err != nil {
		return nil, err
	}
	// remove user from disconnected list if it's in it
	s.Auth.RemoveDisconnectedUser(user.Id)

	response := dto.AuthResponseDTO{
		Need2Fa:      false,
		AccessToken:  at,
		RefreshToken: rt,
	}

	return &response, nil
}

// getRolePermissions returns all user's role permissions map
func getRolePermissions(s services.Permissions, roleID uint64) (map[string]uint8, error) {
	permSearchDTO := dto.SearchPermissionDTO{
		RoleID: roleID,
	}
	perms, err := s.Search(permSearchDTO)
	if err != nil {
		return nil, err
	}
	permissions := make(map[string]uint8)
	for _, perm := range perms {
		permissions[perm.ResourceURIMask] = perm.MethodMask
	}
	return permissions, err
}

// updateUserLastLogin updates user last login timestamp
func updateUserLastLogin(s services.Users, userID uint64) error {
	// update user's last login timestamp
	updateAt := time.Now()
	updateUserDTO := dto.UpdateUserDTO{
		LastLoginAt: &updateAt,
	}
	_, err := s.Update(updateUserDTO)
	if err != nil {
		return err
	}
	return nil
}
