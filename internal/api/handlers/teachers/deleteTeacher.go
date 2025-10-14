package teachers

import (
	"fmt"
	"net/http"
	"school-management-system/pkg/utils"
	"strconv"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid id", http.StatusInternalServerError)
		return

	}
	err = h.service.Delete(r.Context(), id)
	if err != nil {
		if err.Error() == "teacher not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			fmt.Println(err)
			http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
		}
		return
	}

	if err := utils.SendResponse[any](w, r, "teacher deleted successfully", http.StatusOK, nil); err != nil {
		fmt.Println(err)
		http.Error(w, "Error giving response", http.StatusInternalServerError)
		return
	}

}
