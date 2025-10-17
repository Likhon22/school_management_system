package exec

import (
	"net/http"
	"school-management-system/pkg/utils"
)

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {

	utils.InvalidateCookie(w, utils.JWTCookieName)

	if err := utils.SendResponse[any](w, r, "logout successfully", http.StatusCreated, nil); err != nil {
		utils.ErrorHandler(w, err, "Error sending response", http.StatusInternalServerError)
		return
	}

}
