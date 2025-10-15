package teachers

import (
	"net/http"
	"school-management-system/pkg/utils"
	"strconv"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var reqTeacher ReqCreateTeacher
	if err := utils.ReadJson(w, r, &reqTeacher); err != nil {
		utils.ErrorHandler(w, err, "Error updating user", http.StatusInternalServerError)
		return
	}
	if err := utils.ReadJson(w, r, &reqTeacher); err != nil {
		utils.ErrorHandler(w, err, "Error updating user", http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorHandler(w, err, "Invalid id type", http.StatusInternalServerError)
		return
	}

	updateMap := utils.StructToMap(reqTeacher)
	updated, err := h.service.Update(r.Context(), updateMap, allowedFields, id)
	if err != nil {
		utils.ErrorHandler(w, err, "Error updating user", http.StatusInternalServerError)
		return
	}

	if updated == nil {
		utils.ErrorHandler(w, nil, "no teacher found", http.StatusNotFound)
		return
	}
	if err := utils.SendResponse(w, r, "teacher updated successfully", http.StatusOK, updated); err != nil {
		utils.ErrorHandler(w, err, "Error updating user", http.StatusInternalServerError)
		return
	}

}
