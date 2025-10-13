package utils

import (
	"fmt"
	"strings"
)

func BuildFilteredQuery(baseQuery string, filters map[string]string, useLike bool) (string, []interface{}) {

	var args []interface{}
	argCount := 1

	if len(filters) == 0 {
		return baseQuery, args

	}
	var sb strings.Builder
	sb.WriteString(baseQuery)
	sb.WriteString(" WHERE 1=1")

	for field, value := range filters {
		if useLike {
			sb.WriteString(fmt.Sprintf(" AND %s LIKE $%d", field, argCount))
			args = append(args, value)

		} else {
			sb.WriteString(fmt.Sprintf(" AND %s = $%d", field, argCount))
			args = append(args, value)
		}
		argCount++
	}
	return sb.String(), args

}
