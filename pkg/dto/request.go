package dto

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsImportant bool   `json:"is_important"`
	DueDate     string `json:"due_date"`
}

type CreateEventRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	AppointedDate string `json:"appointed_date"`
}

type SignUpRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
