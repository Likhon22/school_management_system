package subjects

import (
	"net/http"
	"school-management-system/pkg/utils"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	filters := utils.BUildFilters(r, params)
	sortOptions := utils.ParseSortQueryOptions(r, allowedSortFields, "created_at DESC")
	subjects, err := h.service.Get(r.Context(), filters, sortOptions)
	if err != nil {

		utils.ErrorHandler(w, err, "Error fetching subjects", http.StatusInternalServerError)
		return
	}
	if err := utils.SendResponse(w, r, "all subjects retrieved successfully", http.StatusCreated, &subjects); err != nil {

		utils.ErrorHandler(w, err, "Error giving response", http.StatusInternalServerError)
		return
	}

}
