package exec

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/pkg/utils"

	"github.com/lib/pq"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var reqExec ReqCreateExec
	if err := utils.ReadJson(w, r, &reqExec); err != nil {
		utils.ErrorHandler(w, err, "something went wrong", http.StatusInternalServerError)
		return
	}

	if err := h.validator.ValidateStruct(reqExec); err != nil {
		utils.ErrorHandler(w, err, "Validation failed", http.StatusBadRequest)
		return
	}

	created, err := h.service.Create(r.Context(), &models.Exec{
		FirstName: reqExec.FirstName,
		LastName:  reqExec.LastName,
		Email:     reqExec.Email,
		Username:  reqExec.Username,
		Password:  reqExec.Password,
		Role:      models.Role(reqExec.Role),
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23505":
				utils.ErrorHandler(w, err, "email or username already exists", http.StatusBadRequest)
				return
			}
		}
		utils.ErrorHandler(w, err, "Error creating exec", http.StatusInternalServerError)
		return
	}

	if err := utils.SendResponse(w, r, "exec created successfully", http.StatusCreated, created); err != nil {
		utils.ErrorHandler(w, err, "Error sending response", http.StatusInternalServerError)
		return
	}
}
