package exec

import (
	"net/http"
	"school-management-system/internal/service"
	"school-management-system/pkg/utils"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var reqLoginExec ReqLoginExec
	if err := utils.ReadJson(w, r, &reqLoginExec); err != nil {
		utils.ErrorHandler(w, err, "something went wrong", http.StatusInternalServerError)
		return
	}
	if err := h.validator.ValidateStruct(reqLoginExec); err != nil {
		utils.ErrorHandler(w, err, "Validation failed", http.StatusBadRequest)
		return
	}

	info, err := h.service.Login(r.Context(), reqLoginExec.Email, reqLoginExec.Password)
	if err != nil {
		switch err {
		case service.ErrExecNotFound, service.ErrPasswordInvalid:
			// Show safe, user-friendly message
			utils.ErrorHandler(w, err, "Invalid email or password", http.StatusUnauthorized)
		default:
			// Internal or unknown error
			utils.ErrorHandler(w, err, "Something went wrong", http.StatusInternalServerError)
		}
		return
	}
	if err := utils.SendResponse(w, r, "login successfully", http.StatusCreated, info); err != nil {
		utils.ErrorHandler(w, err, "Error sending response", http.StatusInternalServerError)
		return
	}

}
