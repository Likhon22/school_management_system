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
		fmt.Println(err)
		http.Error(w, "Error getting pathValue", http.StatusInternalServerError)
		return

	}
	teacher, err := h.service.GetTeacherById(r.Context(), id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting teacher", http.StatusInternalServerError)
		return
	}
	if teacher == nil {
		http.Error(w, "no teacher found", http.StatusNotFound)
		return
	}
	if err := utils.SendResponse(w, r, "teacher retrieved successfully", http.StatusCreated, &teacher); err != nil {
		fmt.Println(err)
		http.Error(w, "Error sending response", http.StatusInternalServerError)
		return
	}

}
