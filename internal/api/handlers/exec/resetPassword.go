package exec

import (
	"net/http"
	"school-management-system/pkg/utils"
)

func (h *Handler) ResetPassword(w http.ResponseWriter, r *http.Request) {

	if err := utils.SendResponse[any](w, r, "you are there boy", http.StatusOK, nil); err != nil {
		utils.ErrorHandler(w, err, "something went wrong", http.StatusInternalServerError)
	}

}
