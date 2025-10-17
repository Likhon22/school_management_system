package exec

import (
	"errors"
	"net/http"
	"school-management-system/internal/api/contextkeys"
	"school-management-system/internal/service"
	"school-management-system/pkg/utils"
	"time"
)

func (h *Handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := ctx.Value(contextkeys.UIdKey).(int)
	if !ok {
		utils.ErrorHandler(w, errors.New("failed about the context value"), "something went wrong", http.StatusInternalServerError)
		return

	}
	var updatePasswordReq UpdatePasswordReqExec
	if err := utils.ReadJson(w, r, &updatePasswordReq); err != nil {
		utils.ErrorHandler(w, err, "something went wrong", http.StatusInternalServerError)
		return
	}
	if err := h.validator.ValidateStruct(updatePasswordReq); err != nil {
		utils.ErrorHandler(w, err, "Validation failed", http.StatusBadRequest)
		return
	}

	token, err := h.service.UpdatePassword(ctx, userId, updatePasswordReq.CurrentPassword, updatePasswordReq.NewPassword)
	if err != nil {
		switch err {
		case service.ErrPasswordInvalid:
			utils.ErrorHandler(w, err, "Password does not match", http.StatusUnauthorized)
		case service.ErrExecNotFound:
			utils.ErrorHandler(w, err, "Invalid user", http.StatusUnauthorized)
		default:
			// Internal or unknown error
			utils.ErrorHandler(w, err, "Something went wrong", http.StatusInternalServerError)
		}
		return
	}
	utils.InvalidateCookie(w, utils.JWTCookieName)
	http.SetCookie(w, &http.Cookie{
		Name:     utils.JWTCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(1 * time.Hour),
	})
	if err := utils.SendResponse[any](w, r, "password updating successful", http.StatusCreated, nil); err != nil {
		utils.ErrorHandler(w, err, "Error sending response", http.StatusInternalServerError)
		return
	}
}
