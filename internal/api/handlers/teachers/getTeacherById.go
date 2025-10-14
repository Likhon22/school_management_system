package teachers

import (
	"fmt"
	"net/http"
	"school-management-system/pkg/utils"
	"strconv"
)

func (h *Handler) GetTeacherById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorHandler(w, err, "Invalid id", http.StatusBadRequest)
		return

	}
	teacher, err := h.service.GetTeacherById(r.Context(), id)
	if err != nil {
		utils.ErrorHandler(w, err, "Error getting teacher", http.StatusInternalServerError)
		return
	}
	if teacher == nil {
		utils.ErrorHandler(w, nil, "no teacher found", http.StatusNotFound)
		return
	}
	if err := utils.SendResponse(w, r, "teacher retrieved successfully", http.StatusCreated, &teacher); err != nil {
		fmt.Println(err)
		utils.ErrorHandler(w, err, "Error sending response", http.StatusInternalServerError)
		return
	}

}
