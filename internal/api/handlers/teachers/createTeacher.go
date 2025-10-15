package teachers

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/pkg/utils"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var reqTeacher ReqCreateTeacher
	if err := utils.ReadJson(w, r, &reqTeacher); err != nil {

		utils.ErrorHandler(w, err, "Error creating teacher", http.StatusInternalServerError)
		return
	}
	if err := h.validator.ValidateStruct(reqTeacher); err != nil {
		utils.ErrorHandler(w, err, "Validation failed", http.StatusBadRequest)
		return
	}

	created, err := h.service.Create(r.Context(), models.Teacher{
		FirstName: reqTeacher.FirstName,
		LastName:  reqTeacher.LastName,
		Email:     reqTeacher.Email,
		Class:     reqTeacher.Class,
		Subject:   reqTeacher.Subject,
	})
	if err != nil {

		utils.ErrorHandler(w, err, "Error creating teacher", http.StatusInternalServerError)
		return
	}
	if err := utils.SendResponse(w, r, "teacher created successfully", http.StatusCreated, created); err != nil {

		utils.ErrorHandler(w, err, "Error creating teacher", http.StatusInternalServerError)
		return
	}

}
