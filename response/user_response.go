package response

import (
	"learning/models"
	"time"
)

type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func ToUserResponse(u *models.User) *UserResponse {
	res := &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	if u.Biodata != nil {
		res.Phone = u.Biodata.Phone
		res.Address = u.Biodata.Address
	}

	return res
}

func ToCreateUserResponse(u *models.User) *CreateUserResponse {
	res := &CreateUserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}

	if u.Biodata != nil {
		res.Phone = u.Biodata.Phone
		res.Address = u.Biodata.Address
	}

	return res
}
