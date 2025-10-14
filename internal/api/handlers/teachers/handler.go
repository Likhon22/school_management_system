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

var params = map[string]string{
	"first_name": "first_name",
	"last_name":  "last_name",
	"subject":    "subject",
	"class":      "class",
}
var allowedSortFields = map[string]bool{
	"first_name": true,
	"last_name":  true,
	"class":      true,
	"subject":    true,
	"created_at": true,
}

var allowedFields = map[string]bool{
	"first_name": true,
	"last_name":  true,
	"email":      true,
	"class":      true,
	"subject":    true,
}
