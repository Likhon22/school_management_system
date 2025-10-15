package class

import (
	"net/http"
	"school-management-system/pkg/utils"
	"strconv"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var reqClass ReqCreateClass
	if err := utils.ReadJson(w, r, &reqClass); err != nil {
		utils.ErrorHandler(w, err, "Error updating class", http.StatusInternalServerError)
		return
	}
	if err := utils.ReadJson(w, r, &reqClass); err != nil {
		utils.ErrorHandler(w, err, "Error updating class", http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorHandler(w, err, "Invalid id type", http.StatusInternalServerError)
		return
	}

	updateMap := utils.StructToMap(reqClass)
	updated, err := h.service.Update(r.Context(), updateMap, allowedFields, id)
	if err != nil {
		utils.ErrorHandler(w, err, "Error updating class", http.StatusInternalServerError)
		return
	}

	if updated == nil {
		utils.ErrorHandler(w, nil, "no student found", http.StatusNotFound)
		return
	}
	if err := utils.SendResponse(w, r, "student updated successfully", http.StatusOK, updated); err != nil {
		utils.ErrorHandler(w, err, "Error updating class", http.StatusInternalServerError)
		return
	}

}
