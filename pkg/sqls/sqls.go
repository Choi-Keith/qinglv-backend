package sqls

import (
	"fmt"
	"strings"
)

func GenSqlike(field string) string {
	return fmt.Sprintf("%s%s%s", "%", field, "%")
}

func HandleSort(sortString string) string {
	sortField := "created_at"
	sort := "DESC"
	if sortString != "" {
		sortStrs := strings.Split(sortString, "|")
		sortField = sortStrs[0]
		sort = strings.ToUpper(sortStrs[1])
	}
	orderBy := fmt.Sprintf("%s %s", sortField, sort)
	return orderBy
}
