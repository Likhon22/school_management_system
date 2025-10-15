package class

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/pkg/utils"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var reqClass ReqCreateClass
	if err := utils.ReadJson(w, r, &reqClass); err != nil {

		utils.ErrorHandler(w, err, "Error creating class", http.StatusInternalServerError)
		return
	}
	if err := h.validator.ValidateStruct(reqClass); err != nil {
		utils.ErrorHandler(w, err, "Validation failed", http.StatusBadRequest)
		return
	}

	created, err := h.service.Create(r.Context(), models.Class{
		Name: reqClass.Name,
	})
	if err != nil {

		utils.ErrorHandler(w, err, "Error creating class", http.StatusInternalServerError)
		return
	}
	if err := utils.SendResponse(w, r, "class created successfully", http.StatusCreated, created); err != nil {

		utils.ErrorHandler(w, err, "Error creating class", http.StatusInternalServerError)
		return
	}

}
