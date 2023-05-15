package dto

type AuthDataDTO struct {
	UserID      uint64
	UserName    string
	Email       string
	Permissions map[string]uint8
}

type AuthRequestDTO struct {
	Login    string `json:"login"` // user email
	Password string `json:"password"`
}

type AuthResponseDTO struct {
	Need2Fa      bool   `json:"need_2fa,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type RefreshRequestDTO struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponseDTO struct {
	AccessToken string `json:"access_token"`
}

type OTPRequestDTO struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	OtpKey   string `json:"otp_key" binding:"required"`
}
