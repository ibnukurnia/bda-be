package requests

type LoginRequest struct {
	Pernr    string `json:"pernr" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *LoginRequest) Messages() map[string]string {
	return map[string]string{
		"Pernr.required":    "Pernr is Required",
		"Password.required": "Password is Required",
	}
}
