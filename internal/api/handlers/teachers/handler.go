package teachers

type Handler struct {
	service TeacherService
}

func NewHandler(TeacherService TeacherService) *Handler {
	return &Handler{
		service: TeacherService,
	}

}

type ReqCreateTeacher struct {
	Email     string `json:"email" db:"email"`
	FirstName string `json:"first_name,omitempty" db:"first_name"`
	LastName  string `json:"last_name,omitempty" db:"last_name"`
	Class     string `json:"class,omitempty" db:"class"`
	Subject   string `json:"subject,omitempty" db:"subject"`
}
