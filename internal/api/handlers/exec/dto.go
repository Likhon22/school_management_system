package exec

type ReqCreateExec struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string `json:"last_name" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required,min=3,max=100"`
	Password  string `json:"password" validate:"required,min=6"`
	Role      string `json:"role" validate:"required,oneof=admin teacher student exec"`
}

type ReqUpdateExec struct {
	FirstName string `json:"first_name,omitempty" validate:"omitempty,min=2,max=100"`
	LastName  string `json:"last_name,omitempty" validate:"omitempty,min=2,max=100"`
	Email     string `json:"email,omitempty" validate:"omitempty,email"`
	Username  string `json:"username,omitempty" validate:"omitempty,min=3,max=100"`
	Password  string `json:"password,omitempty" validate:"omitempty,min=6"`
	Role      string `json:"role,omitempty" validate:"omitempty,oneof=admin teacher student exec"`
}
