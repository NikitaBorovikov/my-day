package model

type User struct {
	ID          int64  `json:"-"`
	UserName    string `json:"user_name" validate:"min=3,max=100"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"min=7,max=255"`
	EncPassword string `json:"-"`
	RegDate     string `json:"reg_date"`
}

type Task struct {
	ID          int64  `json:"-"`
	UserID      int64  `json:"user_id"`
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"max=2000"`
	IsImportant bool   `json:"is_important"`
	DueDate     string `json:"due_date"`
	CreatedDate string `json:"created_date"`
	IsDone      bool   `json:"is_done"`
}

type Event struct {
	ID            int64  `json:"-"`
	UserID        int64  `json:"user_id"`
	Name          string `json:"name" validate:"required,max=100"`
	Description   string `json:"description" validate:"max=2000"`
	AppointedDate string `json:"appointed_date"`
}

type MyDay struct {
	UserID int64   `json:"user_id"`
	Date   string  `json:"date"`
	Tasks  []Task  `json:"tasks"`
	Events []Event `json:"events"`
}

type UserRepository interface {
	SignUp(u *User) error
	SignIn(email, password string) (*User, error)
}

type TaskRepository interface {
	Create(t *Task) error
	GetAll(userID int64) ([]Task, error)
	GetByID(taskID int64) (*Task, error)
	Update(t *Task) error
	Delete(taskID int64) error
}

type EventRepository interface {
	Create(e *Event) error
	GetAll(userID int64) ([]Event, error)
	GetByID(eventID int64) (*Event, error)
	Update(e *Event) error
	Delete(eventID int64) error
}

type MyDayRepository interface {
	Get(userID int64, date string) (*MyDay, error)
}
