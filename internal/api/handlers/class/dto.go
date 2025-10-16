package class

type ReqCreateClass struct {
	Name string `json:"name" validate:"required,min=1,max=50"`
}
