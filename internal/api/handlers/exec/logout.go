package exec

import (
	"net/http"
	"school-management-system/pkg/utils"
	"time"
)

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})

	if err := utils.SendResponse[any](w, r, "logout successfully", http.StatusCreated, nil); err != nil {
		utils.ErrorHandler(w, err, "Error sending response", http.StatusInternalServerError)
		return
	}

}
