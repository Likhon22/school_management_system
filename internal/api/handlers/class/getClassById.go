package class

import (
	"fmt"
	"net/http"
	"school-management-system/pkg/utils"
	"strconv"
)

func (h *Handler) GetClassById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorHandler(w, err, "Invalid id", http.StatusBadRequest)
		return

	}
	fmt.Println(id)
	class, err := h.service.GetClassById(r.Context(), id)
	if err != nil {
		utils.ErrorHandler(w, err, "Error getting class", http.StatusInternalServerError)
		return
	}
	if class == nil {
		utils.ErrorHandler(w, nil, "no class found", http.StatusNotFound)
		return
	}
	if err := utils.SendResponse(w, r, "class retrieved successfully", http.StatusCreated, &class); err != nil {
		fmt.Println(err)
		utils.ErrorHandler(w, err, "Error sending response", http.StatusInternalServerError)
		return
	}

}
