package exec

import (
	"fmt"
	"net/http"
	"school-management-system/pkg/utils"
	"strconv"
)

func (h *Handler) GetExecById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorHandler(w, err, "Invalid id", http.StatusBadRequest)
		return

	}
	exec, err := h.service.GetExecById(r.Context(), id)
	if err != nil {
		utils.ErrorHandler(w, err, "Error getting exec", http.StatusInternalServerError)
		return
	}
	if exec == nil {
		utils.ErrorHandler(w, nil, "no exec found", http.StatusNotFound)
		return
	}
	if err := utils.SendResponse(w, r, "exec retrieved successfully", http.StatusCreated, &exec); err != nil {
		fmt.Println(err)
		utils.ErrorHandler(w, err, "Error sending response", http.StatusInternalServerError)
		return
	}

}
