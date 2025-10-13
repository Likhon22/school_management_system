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
	filters := utils.BUildFilters(r, params)

	fmt.Println(filters)
	teachers, err := h.service.Get(r.Context(), filters)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}
	if err := utils.SendResponse(w, r, "all teachers retrieved successfully", http.StatusCreated, &teachers); err != nil {
		fmt.Println(err)
		http.Error(w, "Error giving response", http.StatusInternalServerError)
		return
	}

}
