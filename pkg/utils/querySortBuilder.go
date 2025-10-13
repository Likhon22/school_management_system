package utils

import "fmt"

func BuildSortQuery(sort SortOption) string {
	if sort.Column == "" {
		return ""

	}
	return fmt.Sprintf(" ORDER BY %s %s", sort.Column, sort.Order)

}
