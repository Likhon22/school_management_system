package exec

import (
	"net/http"
	"school-management-system/internal/service"
	"school-management-system/pkg/utils"
)

func (h *Handler) ForgetPassword(w http.ResponseWriter, r *http.Request) {

	var forgetPasswordReqExec ForgetPasswordReqExec
	if err := utils.ReadJson(w, r, &forgetPasswordReqExec); err != nil {
		utils.ErrorHandler(w, err, "something went wrong", http.StatusInternalServerError)
		return
	}
	if err := h.validator.ValidateStruct(forgetPasswordReqExec); err != nil {
		utils.ErrorHandler(w, err, "validation failed", http.StatusInternalServerError)
		return
	}
	s, err := h.service.ForgetPassword(r.Context(), forgetPasswordReqExec.Email)
	if err != nil {
		switch err {
		case service.ErrExecNotFound:
			utils.ErrorHandler(w, err, "invalid email", http.StatusInternalServerError)
		default:
			utils.ErrorHandler(w, err, "something went wrong", http.StatusInternalServerError)
		}
		return

	}

	if err := utils.SendResponse[any](w, r, s, int(http.StatusOK), nil); err != nil {
		utils.ErrorHandler(w, err, "something went wrong", http.StatusInternalServerError)
	}

}
