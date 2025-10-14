package teachers

type ReqCreateTeacher struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50"`
	Class     string `json:"class" validate:"required,alphanum"`
	Subject   string `json:"subject" validate:"required,min=2,max=100"`
}
