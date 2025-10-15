package teachers

import (
	"fmt"
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/pkg/utils"

	"github.com/lib/pq"
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
		ClassID:   reqTeacher.ClassID,
		Subject:   reqTeacher.Subject,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			utils.ErrorHandler(w, err, fmt.Sprintf("class with id %d does not exist", reqTeacher.ClassID), http.StatusBadRequest)
			return

		}
		utils.ErrorHandler(w, err, "Error creating student", http.StatusInternalServerError)
		return
	}
	if err := utils.SendResponse(w, r, "teacher created successfully", http.StatusCreated, created); err != nil {

		utils.ErrorHandler(w, err, "Error creating teacher", http.StatusInternalServerError)
		return
	}

}
