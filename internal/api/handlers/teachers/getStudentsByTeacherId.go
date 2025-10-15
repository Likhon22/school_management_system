package teachers

import (
	"net/http"
	"school-management-system/pkg/utils"
	"strconv"
)

func (h *Handler) GetStudentsByTeacherID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorHandler(w, err, "Invalid id", http.StatusBadRequest)
		return
	}
	students, err := h.service.GetStudentsByTeacherID(r.Context(), id)
	if err != nil {
		utils.ErrorHandler(w, err, "internal server error", http.StatusInternalServerError)
		return
	}
	if err := utils.SendResponse(w, r, "all teachers retrieved successfully", http.StatusCreated, &students); err != nil {

		utils.ErrorHandler(w, err, "Error giving response", http.StatusInternalServerError)
		return
	}

}
