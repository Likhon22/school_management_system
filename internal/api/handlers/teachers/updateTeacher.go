package teachers

import (
	"fmt"
	"net/http"
	"school-management-system/pkg/utils"
	"strconv"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var reqTeacher ReqCreateTeacher
	if err := utils.ReadJson(w, r, &reqTeacher); err != nil {
		fmt.Println(err)
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid id type", http.StatusInternalServerError)
		return
	}

	updateMap := utils.StructToMap(reqTeacher)
	updated, err := h.service.Update(r.Context(), updateMap, allowedFields, id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	if updated == nil {
		http.Error(w, "no teacher found", http.StatusNotFound)
		return
	}
	if err := utils.SendResponse(w, r, "teacher updated successfully", http.StatusOK, updated); err != nil {
		fmt.Println(err)
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

}
