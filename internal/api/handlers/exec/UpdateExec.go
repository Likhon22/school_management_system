package exec

import (
	"net/http"
	"school-management-system/pkg/utils"
	"strconv"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var reqExec ReqUpdateExec
	if err := utils.ReadJson(w, r, &reqExec); err != nil {
		utils.ErrorHandler(w, err, "Error updating exec", http.StatusInternalServerError)
		return
	}
	if err := utils.ReadJson(w, r, &reqExec); err != nil {
		utils.ErrorHandler(w, err, "Error updating exec", http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorHandler(w, err, "Invalid id type", http.StatusInternalServerError)
		return
	}

	updateMap := utils.StructToMap(reqExec)
	updated, err := h.service.Update(r.Context(), updateMap, allowedFields, id)
	if err != nil {
		utils.ErrorHandler(w, err, "Error updating exec", http.StatusInternalServerError)
		return
	}

	if updated == nil {
		utils.ErrorHandler(w, nil, "no exec found", http.StatusNotFound)
		return
	}
	if err := utils.SendResponse(w, r, "exec updated successfully", http.StatusOK, updated); err != nil {
		utils.ErrorHandler(w, err, "Error updating exec", http.StatusInternalServerError)
		return
	}

}
