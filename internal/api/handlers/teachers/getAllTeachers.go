package teachers

import (
	"fmt"
	"net/http"
	"school-management-system/pkg/utils"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	teachers, err := h.service.Get(r.Context(), firstName, lastName)
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
