package usecases

import (
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/densmart/users-manager/internal/domain/services"
	"github.com/densmart/users-manager/internal/domain/utils"
	"github.com/densmart/users-manager/internal/logger"
	"time"
)

func CreateUser(s services.Service, data dto.CreateUserDTO) (*dto.UserDTO, error) {
	// create password hash
	pwdHash, err := utils.GeneratePasswordHash(data.Password)
	if err != nil {
		return nil, err
	}
	originalPwd := data.Password
	data.Password = pwdHash
	// create random OTP secret
	data.Token2fa, err = utils.Create2FASecret()
	if err != nil {
		return nil, err
	}
	user, err := s.Users.Create(data)
	if err != nil {
		return nil, &dto.APIError{
			HttpCode: 400,
			Message:  err.Error(),
		}
	}
	response := dto.UserDTO{
		ID:        user.Id,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone.String,
		IsActive:  user.IsActive,
		Is2fa:     user.Is2fa,
		RoleID:    user.RoleID,
	}

	// send email with pwd and QR code
	if err = sendGreetEmail(user, originalPwd); err != nil {
		logger.Errorf(err.Error())
	}

	return &response, nil
}

func UpdateUser(s services.Service, data dto.UpdateUserDTO) (*dto.UserDTO, error) {
	user, err := s.Users.Update(data)
	if err != nil {
		return nil, err
	}
	response := dto.UserDTO{
		ID:        user.Id,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone.String,
		IsActive:  user.IsActive,
		Is2fa:     user.Is2fa,
		RoleID:    user.RoleID,
	}
	return &response, nil
}

func RetrieveUser(s services.Service, id uint64) (*dto.UserDTO, error) {
	user, err := s.Users.Retrieve(id)
	if err != nil {
		return nil, err
	}
	response := dto.UserDTO{
		ID:        user.Id,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone.String,
		IsActive:  user.IsActive,
		Is2fa:     user.Is2fa,
		RoleID:    user.RoleID,
	}
	return &response, nil
}

func SearchUsers(s services.Service, data dto.SearchUserDTO) (*dto.UsersDTO, error) {
	paginator := utils.NewPaginator(data.BaseSearchRequestDto)
	offset := paginator.GetOffset()
	limit := paginator.GetLimit()
	data.Offset = &offset
	data.Limit = &limit

	actions, err := s.Users.Search(data)
	if err != nil {
		return nil, err
	}

	response := dto.UsersDTO{
		Pagination: paginator.ToLinkHeader(),
	}
	for _, user := range actions {
		userDTO := dto.UserDTO{
			ID:        user.Id,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone.String,
			IsActive:  user.IsActive,
			Is2fa:     user.Is2fa,
			RoleID:    user.RoleID,
		}
		response.Items = append(response.Items, userDTO)
	}

	return &response, nil
}

func DeleteUser(s services.Service, id uint64) error {
	return s.Users.Delete(id)
}

// sendGreetEmail Make and send email with password data and QR code and send it to user
func sendGreetEmail(user entities.User, originalPwd string) error {
	// create QR for 2FA application
	otpURL, err := utils.Create2faURL(user.Email, user.Token2fa.String)
	if err != nil {
		return err
	}
	qrData, err := utils.CreateQRCode(otpURL)
	if err != nil {
		return err
	}

	// send email to user
	// make mail body from template
	emailData := map[string]interface{}{
		"username": user.FirstName + " " + user.LastName,
		"email":    user.Email,
		"password": originalPwd,
	}
	mailBody, _ := utils.ParseMailTemplate("welcome", emailData)
	// make mail message
	mail := &utils.MailMessage{
		Subject:     "Welcome to SB",
		Body:        mailBody,
		Attachments: make(map[string][]byte),
	}
	mail.To = []string{user.Email}
	// attach QR to email
	mail.Attach("qr.png", qrData)
	// send email to user
	// @TODO: add tests, check on real mail send
	// mailer initialize
	//mailer := utils.NewMailer()
	//err = mailer.Send(mail)
	//if err != nil {
	//	// if email isn`t send - log error and continue
	//	// user has to be created anyway
	//	return err
	//}
	return nil
}
