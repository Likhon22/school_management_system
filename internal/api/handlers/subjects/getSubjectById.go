package subjects

import (
	"fmt"
	"net/http"
	"school-management-system/pkg/utils"
	"strconv"
)

func (h *Handler) GetSubjectById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorHandler(w, err, "Invalid id", http.StatusBadRequest)
		return

	}
	subject, err := h.service.GetSubjectById(r.Context(), id)
	if err != nil {
		utils.ErrorHandler(w, err, "Error getting subject", http.StatusInternalServerError)
		return
	}
	if subject == nil {
		utils.ErrorHandler(w, nil, "no subject found", http.StatusNotFound)
		return
	}
	if err := utils.SendResponse(w, r, "subject retrieved successfully", http.StatusCreated, &subject); err != nil {
		fmt.Println(err)
		utils.ErrorHandler(w, err, "Error sending response", http.StatusInternalServerError)
		return
	}

}
