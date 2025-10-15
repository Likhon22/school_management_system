package students

import (
	"fmt"
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/pkg/utils"

	"github.com/lib/pq"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var reqStudent ReqCreateStudent
	if err := utils.ReadJson(w, r, &reqStudent); err != nil {

		utils.ErrorHandler(w, err, "Error creating student", http.StatusInternalServerError)
		return
	}
	if err := h.validator.ValidateStruct(reqStudent); err != nil {
		utils.ErrorHandler(w, err, "Validation failed", http.StatusBadRequest)
		return
	}

	created, err := h.service.Create(r.Context(), models.Student{
		FirstName: reqStudent.FirstName,
		LastName:  reqStudent.LastName,
		Email:     reqStudent.Email,
		ClassID:   reqStudent.ClassID,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			utils.ErrorHandler(w, err, fmt.Sprintf("class with id %d does not exist", reqStudent.ClassID), http.StatusBadRequest)
			return

		}
		utils.ErrorHandler(w, err, "Error creating student", http.StatusInternalServerError)
		return
	}

	if err := utils.SendResponse(w, r, "student created successfully", http.StatusCreated, created); err != nil {

		utils.ErrorHandler(w, err, "Error creating student", http.StatusInternalServerError)
		return
	}

}
