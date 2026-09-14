package handler

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authResponse struct {
	Token   string `json:"token,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
