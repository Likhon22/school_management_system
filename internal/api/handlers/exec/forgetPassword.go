package exec

import (
	"fmt"
	"net/http"
	"school-management-system/internal/service"
	"school-management-system/pkg/utils"
)

func (h *Handler) ForgetPassword(w http.ResponseWriter, r *http.Request) {

	var email string
	if err := utils.ReadJson(w, r, &email); err != nil {
		utils.ErrorHandler(w, err, "something went wrong", http.StatusInternalServerError)
		return
	}
	s, err := h.service.ForgetPassword(r.Context(), email)
	switch err {
	case service.ErrExecNotFound:
		utils.ErrorHandler(w, err, "invalid email", http.StatusInternalServerError)
	default:
		utils.ErrorHandler(w, err, "something went wrong", http.StatusInternalServerError)
	}
	fmt.Println(s)

}
