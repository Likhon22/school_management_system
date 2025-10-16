package exec

import (
	"fmt"
	"net/http"
	"school-management-system/internal/api/contextkeys"
	"school-management-system/pkg/utils"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	userId := ctx.Value(contextkeys.UIdKey)
	email := ctx.Value(contextkeys.EmailKey)
	role := ctx.Value(contextkeys.RoleKey)
	username := ctx.Value(contextkeys.UserKey)
	fmt.Println("id", userId, "email", email, "role", role, "username", username)
	filters := utils.BUildFilters(r, params)
	sortOptions := utils.ParseSortQueryOptions(r, allowedSortFields, "created_at DESC")
	execs, err := h.service.Get(ctx, filters, sortOptions)
	if err != nil {

		utils.ErrorHandler(w, err, "Error fetching exec", http.StatusInternalServerError)
		return
	}
	if err := utils.SendResponse(w, r, "all execs retrieved successfully", http.StatusCreated, &execs); err != nil {

		utils.ErrorHandler(w, err, "Error giving response", http.StatusInternalServerError)
		return
	}

}
