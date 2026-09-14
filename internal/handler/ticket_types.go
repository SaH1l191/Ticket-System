package handler

type createTicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ticketResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ticketErrorResponse struct {
	Error string `json:"error"`
}

type updateStatusRequest struct {
	Status string `json:"status"`
}
