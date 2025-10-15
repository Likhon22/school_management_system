package students

import (
	"fmt"
	"net/http"
	"school-management-system/pkg/utils"
	"strconv"
)

func (h *Handler) GetStudentById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorHandler(w, err, "Invalid id", http.StatusBadRequest)
		return

	}
	student, err := h.service.GetStudentById(r.Context(), id)
	if err != nil {
		utils.ErrorHandler(w, err, "Error getting student", http.StatusInternalServerError)
		return
	}
	if student == nil {
		utils.ErrorHandler(w, nil, "no student found", http.StatusNotFound)
		return
	}
	if err := utils.SendResponse(w, r, "student retrieved successfully", http.StatusCreated, &student); err != nil {
		fmt.Println(err)
		utils.ErrorHandler(w, err, "Error sending response", http.StatusInternalServerError)
		return
	}

}
