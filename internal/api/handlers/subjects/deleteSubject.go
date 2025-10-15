package subjects

import (
	"net/http"
	"school-management-system/pkg/utils"
	"strconv"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {

		utils.ErrorHandler(w, err, "Invalid id", http.StatusInternalServerError)
		return

	}
	err = h.service.Delete(r.Context(), id)
	if err != nil {
		if err.Error() == "subject not found" {
			utils.ErrorHandler(w, nil, err.Error(), http.StatusNotFound)
		} else {

			utils.ErrorHandler(w, err, "Error deleting subject", http.StatusInternalServerError)
		}
		return
	}

	if err := utils.SendResponse[any](w, r, "subject deleted successfully", http.StatusOK, nil); err != nil {

		utils.ErrorHandler(w, err, "Error giving response", http.StatusInternalServerError)
		return
	}

}
