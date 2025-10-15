package class

import (
	"net/http"
	"school-management-system/pkg/utils"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	filters := utils.BUildFilters(r, params)
	sortOptions := utils.ParseSortQueryOptions(r, allowedSortFields, "created_at DESC")
	class, err := h.service.Get(r.Context(), filters, sortOptions)
	if err != nil {

		utils.ErrorHandler(w, err, "Error fetching class", http.StatusInternalServerError)
		return
	}
	if err := utils.SendResponse(w, r, "all class retrieved successfully", http.StatusCreated, &class); err != nil {

		utils.ErrorHandler(w, err, "Error giving response", http.StatusInternalServerError)
		return
	}

}
