package utils

import (
	"net/http"
	"strings"
)

type SortOption struct {
	Column string
	Order  string
}

func ParseSortQueryOptions(r *http.Request, allowedFields map[string]bool, defaultSort string) SortOption {
	sortBy := r.URL.Query().Get("sort_by")
	order := strings.ToUpper(r.URL.Query().Get("order"))
	if !allowedFields[sortBy] {
		sortBy = ""
	}
	if order != "ASC" && order != "DESC" {
		order = "ASC"
	}
	if sortBy == "" && defaultSort != "" {
		parts := strings.Fields(defaultSort)
		if len(parts) == 2 && allowedFields[parts[0]] {
			sortBy = parts[0]
			tmpOrder := strings.ToUpper(parts[1])
			if tmpOrder == "ASC" || tmpOrder == "DESC" {
				order = tmpOrder
			} else {
				order = "ASC"
			}
		} else {
			sortBy = ""
			order = "ASC"
		}
	}

	return SortOption{
		Column: sortBy,
		Order:  order,
	}

}
