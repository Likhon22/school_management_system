package teachers

import (
	"fmt"
	"net/http"
	"school-management-system/pkg/utils"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"subject":    "subject",
		"class":      "class",
	}
	allowedSortFields := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"class":      true,
		"subject":    true,
		"created_at": true,
	}
	filters := utils.BUildFilters(r, params)
	sortOptions := utils.ParseSortQueryOptions(r, allowedSortFields, "created_at DESC")
	teachers, err := h.service.Get(r.Context(), filters, sortOptions)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error fetching teacher", http.StatusInternalServerError)
		return
	}
	if err := utils.SendResponse(w, r, "all teachers retrieved successfully", http.StatusCreated, &teachers); err != nil {
		fmt.Println(err)
		http.Error(w, "Error giving response", http.StatusInternalServerError)
		return
	}

}
