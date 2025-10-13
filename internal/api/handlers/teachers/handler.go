package teachers

import (
	"fmt"
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/pkg/utils"
)

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

func (h *Handler) CREATE(w http.ResponseWriter, r *http.Request) {
	var reqTeacher ReqCreateTeacher
	if err := utils.ReadJson(w, r, &reqTeacher); err != nil {
		fmt.Println(err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	created, err := h.service.CREATE(r.Context(), models.Teacher{
		FirstName: reqTeacher.FirstName,
		LastName:  reqTeacher.LastName,
		Email:     reqTeacher.Email,
		Class:     reqTeacher.Class,
		Subject:   reqTeacher.Subject,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	if err := utils.SendResponse(w, r, "teacher created successfully", http.StatusCreated, created); err != nil {
		fmt.Println(err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

}
