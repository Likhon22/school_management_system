package exec

import (
	"net/http"
	"school-management-system/internal/service"
	"school-management-system/pkg/utils"
)

func (h *Handler) ResetPassword(w http.ResponseWriter, r *http.Request) {

	hashToken := r.PathValue("resetcode")
	var resetPasswordReqExec ResetPasswordReqExec
	if err := utils.ReadJson(w, r, &resetPasswordReqExec); err != nil {
		utils.ErrorHandler(w, err, "something went wrong", http.StatusInternalServerError)
		return
	}
	if err := h.validator.ValidateStruct(resetPasswordReqExec); err != nil {
		utils.ErrorHandler(w, err, "validation failed", http.StatusInternalServerError)
		return
	}
	message, err := h.service.ResetPassword(r.Context(), hashToken, resetPasswordReqExec.NewPassword, resetPasswordReqExec.ConfirmPassword)
	if err != nil {
		switch err {
		case service.ErrExecNotFound:
			utils.ErrorHandler(w, err, "Invalid user", http.StatusUnauthorized)
		case service.ErrPasswordDoestNotMatch:
			utils.ErrorHandler(w, err, "Password does not match", http.StatusBadRequest)

		case service.ErrTokenExpire:
			utils.ErrorHandler(w, err, "Invalid token or token time expired", http.StatusForbidden)
		default:
			utils.ErrorHandler(w, err, "something went wrong", http.StatusInternalServerError)
		}

		return
	}
	if err := utils.SendResponse[any](w, r, message, http.StatusOK, nil); err != nil {
		utils.ErrorHandler(w, err, "something went wrong", http.StatusInternalServerError)
		return
	}

}
