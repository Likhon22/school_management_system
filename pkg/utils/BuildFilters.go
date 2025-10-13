package utils

import "net/http"

func BUildFilters(r *http.Request, paramMap map[string]string) map[string]string {
	filterMap := make(map[string]string)
	for param, dbField := range paramMap {

		value := r.URL.Query().Get(param)
		if value != "" {
			filterMap[dbField] = value

		}

	}
	return filterMap

}
