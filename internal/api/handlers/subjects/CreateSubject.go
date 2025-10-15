package subjects

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/pkg/utils"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var reqSubject ReqCreateSubject
	if err := utils.ReadJson(w, r, &reqSubject); err != nil {

		utils.ErrorHandler(w, err, "Error creating subject", http.StatusInternalServerError)
		return
	}
	if err := h.validator.ValidateStruct(reqSubject); err != nil {
		utils.ErrorHandler(w, err, "Validation failed", http.StatusBadRequest)
		return
	}

	created, err := h.service.Create(r.Context(), models.Subject{
		Name: reqSubject.Name,
	})
	if err != nil {

		utils.ErrorHandler(w, err, "Error creating subject", http.StatusInternalServerError)
		return
	}
	if err := utils.SendResponse(w, r, "subject created successfully", http.StatusCreated, created); err != nil {

		utils.ErrorHandler(w, err, "Error creating subject", http.StatusInternalServerError)
		return
	}

}
